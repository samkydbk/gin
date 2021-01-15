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
	router.GET("/categories", controllers.ListCategory)
	router.GET("/categories/new", controllers.NewCategory)
	router.POST("/categories/create", controllers.CreateCategory)
	router.GET("/categories/edit/:id", controllers.EditCategory)
	router.POST("/categories/update/:id", controllers.UpdateCategory)
	router.GET("/categories/delete/:id", controllers.DeleteCategory)
	router.GET("/posts", controllers.ListPost)
	router.GET("/posts/new", controllers.NewPost)
	router.POST("/posts/create", controllers.CreatePost)
	router.GET("/posts/edit/:id", controllers.EditPost)
	router.POST("/posts/update/:id", controllers.UpdatePost)
	router.GET("/posts/delete/:id", controllers.DeletePost)
	router.GET("/user/list", controllers.ListUser)
	return router
}
