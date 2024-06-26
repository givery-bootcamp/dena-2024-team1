package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
