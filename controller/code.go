package controller

const (
	SUCCESS = 200
	ERROR   = 500
	//约定状态码
	//code=1000...用户模块错误
	CODE_USERNAME_USED  = 1001
	CODE_PASSWORD_WRONG = 1002
	CODE_USER_NOT_EXIST = 1003

	CODE_TOKEN_NOT_EXIST = 1004
	CODE_TOKEN_OUT_TIME  = 1005
	CODE_TOKEN_WRONG     = 1006
	CODE_TYPE_WRONG      = 1007
	CODE_PARAM_WRONG     = 1008

	//code=2000...文章模块错误

	//code=3000...分类模块错误
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	CODE_USERNAME_USED:   "该用户名已存在",
	CODE_PASSWORD_WRONG:  "密码错误",
	CODE_USER_NOT_EXIST:  "用户不存在",
	CODE_TOKEN_NOT_EXIST: "token不存在",
	CODE_TOKEN_OUT_TIME:  "token已过期",
	CODE_TOKEN_WRONG:     "token错误",
	CODE_TYPE_WRONG:      " token格式错误",
	CODE_PARAM_WRONG:     "参数错误",
}

func GetCodeMsg(code int) string {
	return codeMsg[code]
}
