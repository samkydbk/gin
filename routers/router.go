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
	//post请求
	//router.POST("/user/search", controllers.SearchUser)
	//get请求方式一  127.0.0.1:8081/user/search/5  seo友好请求方式
	//router.GET("/user/search/:id", controllers.SearchUser)
	//get请求方式二  127.0.0.1:8081/user/search?id=8
	//router.GET("/user/search", controllers.SearchUser)

	router.GET("/user/search", controllers.SearchUser)

	router.POST("/user/update", controllers.UpdateUser)
	router.GET("/user/del/:id", controllers.DelUser)
	return router
}
