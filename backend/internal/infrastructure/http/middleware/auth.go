package middleware

import (
	"myapp/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func Auth(sessionStore *sessions.CookieStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := sessionStore.Get(ctx.Request, config.SessionName)
		if session.Values[config.SessionKey] == nil {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
