package routers

import (
	"myweb/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/error", controllers.ErrorPage)
	//用户增删改查及列表
	router.GET("/user/list", controllers.UserList)
	router.POST("/user/create", controllers.createUser)
	router.POST("/user/search/:id", controllers.searchUser)
	router.GET("/user/del/:id", controllers.delUser)
	return router
}
