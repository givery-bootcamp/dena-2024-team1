package middleware

import (
	myConfig "myapp/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{myConfig.CorsAllowOrigin}
	config.AllowCredentials = true
	defaultAllowHeaders := cors.DefaultConfig().AllowHeaders
	// X-Custom-Header-For-Preflightを追加
	config.AllowHeaders = append(defaultAllowHeaders, "X-Custom-Header-For-Preflight")
	return cors.New(config)
}
