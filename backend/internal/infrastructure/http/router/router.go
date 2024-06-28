package router

import (
	"myapp/internal/infrastructure/http/middleware"
	"myapp/internal/registry"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.Session())

	setupEndpoints(router)

	return router
}

func setupEndpoints(router *gin.Engine) {
	apiHandler := registry.NewAPIHandler()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	// 認証が不要なエンドポイント
	router.POST("/signup", apiHandler.UserHandler.Signup)
	router.POST("/signin", apiHandler.UserHandler.Signin)
	router.GET("/sketches", apiHandler.SketchHandler.GetSketches)

	authMiddleware := middleware.Auth()

	// 認証が必要なエンドポイント
	router.GET("/hello", authMiddleware, apiHandler.HelloWorldHandler.HelloWorld)
	router.GET("/posts", authMiddleware, apiHandler.PostHandler.GetPosts)
	router.GET("/posts/:id", authMiddleware, apiHandler.PostHandler.GetPost)
	router.POST("/posts", authMiddleware, apiHandler.PostHandler.CreatePost)
	router.POST("/posts/sketches", authMiddleware, apiHandler.SketchHandler.CreateSketch)
	router.PUT("/posts/:id", authMiddleware, apiHandler.PostHandler.UpdatePost)
	router.DELETE("/posts/:id", authMiddleware, apiHandler.PostHandler.DeletePost)
	router.POST("/signout", authMiddleware, apiHandler.UserHandler.Signout)
	router.GET("/session_user", authMiddleware, apiHandler.UserHandler.GetSessionUser)
	router.POST("/sketches", authMiddleware, apiHandler.SketchHandler.CreateSketch)
	router.DELETE("/sketches/:id", authMiddleware, apiHandler.SketchHandler.DeleteSketch)
}
