package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/project-user/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

func (*RouterUser) Route(r *gin.Engine) {
	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
