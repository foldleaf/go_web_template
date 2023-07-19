package mysql

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"web_app/models"

	"go.uber.org/zap"
	"golang.org/x/crypto/scrypt"
)

// 把每一步数据库操作封装成函数
// 待 logic 层调用

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	// 在用户表中查找 username对应的 id，计算数量大于 0 则已存在
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条新的记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = ScryptPassword(user.Password)
	// 执行 SQL 语句入库
	sqlStr := `insert into user (user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10 //长度
	salt := make([]byte, 8)

	salt = []byte{12, 22, 111, 46, 82, 3, 7, 21} // 加盐
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		zap.L().Error(err.Error())
	}
	finalPw := base64.StdEncoding.EncodeToString(HashPw)
	return finalPw
}

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录时输入的密码
	sqlStr := `select username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	fmt.Println("1")
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		fmt.Println("2")
		return err
	}
	// 判断密码是否正确
	password := ScryptPassword(oPassword) // 用户输入的密码加密后与数据库中的密码比对
	if password != user.Password {
		fmt.Println("3")
		return errors.New("用户密码错误")
	}
	fmt.Println("4")

	return
}
