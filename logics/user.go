package logics

import (
	"myweb/models"
)

//用户列表
func UserList() ([]models.User, error) {
	var user models.User
	result := db.Order("id desc").Find(&user)
	return user, result.Error
}

//增加用户
func UserCreate() (int64, error) {
	user := models.User{Name: "abc", Sex: 1, Phone: "13228016321"}
	result := db.Create(&user) // 通过数据的指针来创建
	return user.ID, result.Error
}

//删除用户

//查询用户

//修改用户
