package router

import (
	"myapp/internal/controller/api"
	"myapp/internal/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Transaction())
	router.Use(middleware.Cors())

	SetupEndpoints(router)

	return router
}

func SetupEndpoints(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	router.GET("/hello", api.HelloWorld)
	router.GET("/posts", api.GetPosts)
	router.GET("/posts/:id", api.GetPost)
}
