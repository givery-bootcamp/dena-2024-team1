package router

import (
	"myapp/internal/controller/api"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", api.HelloWorld)
	app.GET("/posts", api.GetPosts)
	app.GET("/posts/:id", api.GetPost)
}
