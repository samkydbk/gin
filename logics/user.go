package logics

import (
	"myweb/models"
)

//用户列表
func UserList(datas []models.User, page int) ([]models.User, int64, error) {
	var pageSize = 2
	offset := (page - 1) * pageSize
	result := db.Order("id desc").Offset(offset).Limit(pageSize).Find(&datas)
	return datas, result.RowsAffected, result.Error
}

//增加用户
func createUser(data models.User) (int64, error) {
	result := db.Create(&data)
	return data.ID, result.Error
	// user := models.User{Name: "abc", Sex: 1, Phone: "13228016321"}
	// result := db.Create(&user) // 通过数据的指针来创建
	// return user.ID, result.Error
}

//查询用户
func searchUser(id int64) (models.User, error) {
	var model models.User
	result := db.First(&model, id)
	return model, result.Error
}

//修改用户
func updateUser(data models.User, id int64) (int64, error) {
	var model models.User
	row := db.First(&model, id)
	if row.Error == nill {
		result := db.Model(&model).Updates(&data)
		return model.ID, result.Error
	}
	return 0, row.Error
}

//删除用户
func delUser(id int64) (int64, error) {
	var model models.User
	result := db.Delete(&model, id)
	return result.RowsAffected, result.Error
}
