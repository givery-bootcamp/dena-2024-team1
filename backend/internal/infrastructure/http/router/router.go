package router

import (
	"myapp/internal/infrastructure/database"
	"myapp/internal/infrastructure/http/middleware"
	"myapp/internal/infrastructure/sessionstore"
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
	sessionStore := sessionstore.GetStore()
	authMiddleware := middleware.Auth(sessionStore)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	// 認証が不要なエンドポイント
	router.POST("/signup", apiHandler.UserHandler.Signup)
	router.POST("/signin", apiHandler.UserHandler.Signin)

	// 認証が必要なエンドポイント
	router.GET("/hello", authMiddleware, apiHandler.HelloWorldHandler.HelloWorld)
	router.GET("/posts", authMiddleware, apiHandler.PostHandler.GetPosts)
	router.GET("/posts/:id", authMiddleware, apiHandler.PostHandler.GetPost)
	router.POST("/posts", authMiddleware, apiHandler.PostHandler.CreatePost)
	router.POST("/posts/sketches", authMiddleware, apiHandler.PostHandler.CreateSketch)
	router.PUT("/posts/:id", authMiddleware, apiHandler.PostHandler.UpdatePost)
	router.DELETE("/posts/:id", authMiddleware, apiHandler.PostHandler.DeletePost)
	router.POST("/signout", authMiddleware, apiHandler.UserHandler.Signout)
	router.GET("/session_user", authMiddleware, apiHandler.UserHandler.GetSessionUser)
	router.GET("/sketches", authMiddleware, apiHandler.SketchHandler.GetSketches)
}
