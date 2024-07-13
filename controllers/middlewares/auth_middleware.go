package middlewares

import (
	"github.com/gin-gonic/gin"
	"gobase/config"
	"gobase/utils"
	"net/http"
	"strings"
)

func Authenticate(ctx *gin.Context) {
	// Get JWT token from ctx
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var token string
	// get token from header
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	// Verify token
	claims, err := utils.ExtractClaimsFromToken(token, config.GetConfig().Security.JWTSecret)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	ctx.Set("claims", claims)
}
