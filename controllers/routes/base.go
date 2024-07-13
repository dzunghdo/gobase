package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBaseRoutes(router *gin.RouterGroup) {
	router.GET("/health", healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
