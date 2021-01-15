package logics

import (
	"myweb/models"
)

func ListUser(datas []models.User) ([]models.User, error) {
	result := db.Order("id desc").Find(&datas)
	return datas, result.Error
}
