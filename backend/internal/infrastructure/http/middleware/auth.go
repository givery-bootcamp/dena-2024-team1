package middleware

import (
	"myapp/internal/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ヘッダーからX-Custom-Header-For-Preflightを取得（フロントエンドがPreflightリクエストを強制するために設定しているヘッダー）
		xCustomHeaderForPreflight := ctx.GetHeader("X-Custom-Header-For-Preflight")

		// X-Custom-Header-For-Preflightが設定されていない場合はエラーを返す
		if xCustomHeaderForPreflight != "true" {
			ctx.JSON(401, gin.H{"error": "Not Set X-Custom-Header-For-Preflight"})
			ctx.Abort()
			return
		}

		session := sessions.Default(ctx)

		if session.Get(config.SessionKey) == nil {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
