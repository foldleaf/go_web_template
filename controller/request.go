package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUsername = "username"

func GetCurrentUser(c *gin.Context) (currentUsername string, err error) {
	username, ok := c.Get(CtxUsername)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	currentUsername, ok = username.(string)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	return currentUsername, nil
}
