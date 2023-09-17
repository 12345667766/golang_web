package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 优雅启停
func Run(r *gin.Engine, srvName string, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	fmt.Println(s)
	log.Printf("Shutting Down project %s...\n", srvName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 最后关闭定时器
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown, cause by : ", srvName, err)
	}
	log.Printf("%s stop success...\n", srvName)
}
