package routes

import (
	"github.com/gin-gonic/gin"
	"gobase/controllers/handlers"
)

func HandleAuthRoutes(router *gin.RouterGroup) {
	authHandler := handlers.NewAuthHandler()
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
}
