package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoverApp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			}
		}()
		ctx.Next()
	}
}
