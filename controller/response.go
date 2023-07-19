package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 1001	// 错误码
	"msg": "xxx"	// 提示信息
	"data": "xxx"	// 数据
}
*/

type ResponseData struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseErr(c *gin.Context, code int) {
	rd := &ResponseData{
		Code: code,
		Msg:  GetCodeMsg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: SUCCESS,
		Msg:  GetCodeMsg(SUCCESS),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrWithMsg(c *gin.Context, code int, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}
