package middleware

import (
	"myapp/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func Transaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.DB.Begin()
		defer func() {
			if 400 <= ctx.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		ctx.Set("db", db)
		ctx.Next()
	}
}
