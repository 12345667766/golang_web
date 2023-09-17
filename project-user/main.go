package main

import (
	"github.com/gin-gonic/gin"
	"log"
	common "test.com/project-common"
	"test.com/project-common/config"
	"test.com/project-common/logs"
	_ "test.com/project-user/api"
	"test.com/project-user/router"
)

func main() {
	config.InitConfig()

	lc := &logs.LogConfig{
		DebugFileName: config.C.LogConfig.DebugFileName,
		InfoFileName:  config.C.LogConfig.InfoFileName,
		WarnFileName:  config.C.LogConfig.WarnFileName,
		MaxSize:       config.C.LogConfig.MaxSize,
		MaxAge:        config.C.LogConfig.MaxAge,
		MaxBackups:    config.C.LogConfig.MaxBackups,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	router.InitRouter(r)
	common.Run(r, config.C.ServerConfig.ServerName, config.C.ServerConfig.Address)
}
