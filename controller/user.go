package controller

import (
	"fmt"
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数，参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		// 记录错误日志
		fmt.Println("请求参数错误:", err)
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(p)
	// 手动对参数进行详细的业务规则的校验
	// 用户名、密码、重复密码不能为空，密码与重复密码要相同
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SignUp with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请正确填写信息",
	// 	})
	// 	return
	// }

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		fmt.Println("请求参数错误:", err)
		zap.L().Error("Login with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 业务处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed ", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "登录成功",
	// })
}
