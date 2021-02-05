package controllers

import (
	"fmt"
	"myweb/logics"
	"myweb/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户列表
func UserList(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	if page == 0 {
		page = 1
	}
	var list []models.User
	res, rows, err := logics.UserList(list, page)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "user.html", gin.H{"title": "用户列表", "list": res, "rows": rows})
}

//创建用户
func CreateUser(c *gin.Context) {
	var model models.User
	model.Name = "eeeee"
	model.Sex = 1
	model.Phone = "13228016321"
	_, err := logics.CreateUser(model)
	if err != nil {
		panic(err)
	}
	fmt.Println("添加成功")
}

//查询用户
func SearchUser(c *gin.Context) {
	//获取post表单信息
	//id := c.PostForm("id")
	//fmt.Println("id:", id)

	//获取get信息 请求方式一
	//id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	//fmt.Println("id:", c.Param("id"))

	//获取get信息 请求方式二
	//fmt.Println("id:", c.Query("id"))

	//将字符串转换为数字函数ParseInt 字符串,进制,返回大小
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	res, err := logics.SearchUser(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("查询成功")
	fmt.Println(res)
}

//修改用户
func UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	var model models.User
	model.Phone = "13288888888"
	_, err := logics.UpdateUser(model, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("修改成功")
}

//删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err := logics.DelUser(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("删除成功")
}
