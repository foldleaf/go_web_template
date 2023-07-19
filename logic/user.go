package logic

import (
	// "errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 生成 UID
	userID := snowflake.GenID()
	// 构造user实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 密码加密
	// 保存进数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		// UserID:   0,
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成 jwt token

	return jwt.GenToken(user.UserID, user.Username)

}
