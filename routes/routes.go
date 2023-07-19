package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signUp", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		// 如果用户登录，判断请求头是否有有效的 jwt token
		ctx.Request.Header.Get("Authorization")
		isLogin := true
		if isLogin {
			ctx.String(http.StatusOK, "pong")
		} else {
			ctx.String(http.StatusOK, "请登录")
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return r
}
