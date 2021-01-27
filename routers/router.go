package routers

import (
	"myweb/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")
	//用户增删改查及列表
	router.GET("/user/list", controllers.UserList)
	router.POST("/user/create", controllers.CreateUser)
	router.POST("/user/search/:id", controllers.SearchUser)
	router.POST("/user/update/:id", controllers.UpdateUser)
	router.GET("/user/del/:id", controllers.DelUser)
	return router
}
