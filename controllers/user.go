package controllers

import (
	"fmt"
	"myweb/logics"
	"myweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var list []models.User
	res, err := logics.ListUser(list)
	if err != nil {
		fmt.Println("发生错误")
	}
	c.HTML(http.StatusOK, "user.html", gin.H{"title": "用户列表", "list": res})
}
