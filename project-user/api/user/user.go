package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	common "test.com/project-common"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		// 返回了接口的实现
		cache: dao.Rc,
	}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	rsp := &common.Result{}
	// 1. 获取参数
	mobile := ctx.PostForm("mobile")
	if !common.VerifyMobile(mobile) {
		// 避免魔法数字
		ctx.JSON(http.StatusOK, rsp.Faile(model.NoLegalMobile, "手机号不合法"))
		return
	}
	code := "123456"
	go func() {
		zap.L().Info("短信平台调用成功，发送短信 INFO")
		zap.L().Debug("短信平台调用成功，发送短信 Debug")
		zap.L().Error("短信平台调用成功，发送短信 Error")
		h := New()
		time.Sleep(2)
		log.Println("短信平台调用成功，发送短信")
		// 存储redis
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			fmt.Println("存储redis失败: ", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, code)
		redisStr, _ := h.cache.Get(c, "REGISTER_"+mobile)
		fmt.Println("redis存入: ", redisStr)
	}()

	// 2. 校验参数
	// 3. 生成验证码 (随机4位或者6位)
	// 4. 调用短信平台（三方平台 放入goroutine, 接口可以快速响应）
	// 5. 存储验证码 存入redis中，过期时间15min
	ctx.JSON(200, rsp.Success(code))
}
