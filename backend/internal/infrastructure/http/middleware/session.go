package middleware

import (
	"myapp/internal/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	store := cookie.NewStore([]byte(config.SessionSecret))

	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		Secure:   true,
	})

	return sessions.Sessions(config.SessionName, store)
}
