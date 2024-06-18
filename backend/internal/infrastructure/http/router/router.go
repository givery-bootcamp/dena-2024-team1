package router

import (
	"myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/http/middleware"
	"myapp/internal/registry"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	db := database.SetupDB()

	router := gin.Default()
	router.Use(middleware.Transaction(db))
	router.Use(middleware.Cors())

	setupEndpoints(router)

	return router
}

func setupEndpoints(router *gin.Engine) {
	apiHandler := registry.NewAPIHandler()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	router.GET("/hello", apiHandler.HelloWorldHandler.HelloWorld)
	router.GET("/posts", apiHandler.PostHandler.GetPosts)
	router.GET("/posts/:id", apiHandler.PostHandler.GetPost)
	router.POST("/posts", apiHandler.PostHandler.CreatePost)
	router.DELETE("/posts/:id", apiHandler.PostHandler.DeletePost)
}
