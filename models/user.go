package models

import "time"

//定义user表模型 结构体
//gorm官方对model说明 https://gorm.io/zh_CN/docs/models.html
type User struct {
	ID        uint64
	Name      string
	Sex       uint64
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
