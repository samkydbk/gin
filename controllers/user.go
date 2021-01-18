package controllers

import (
	"fmt"
	"myweb/logics"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户列表
func UserList(c *gin.Context) {
	//var list []models.User
	//res, err := logics.ListUser(list)
	res, err := logics.UserList()
	if err != nil {
		fmt.Println("发生错误")
	}
	c.HTML(http.StatusOK, "user.html", gin.H{"title": "用户列表", "list": res})
}

//创建用户
func UserCreate(c *gin.Context) {
	res, err := logics.UserCreate()
	if err != nil {
		panic(err)
	}
	fmt.Println(888)
	fmt.Println(res)
}
