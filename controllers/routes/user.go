package routes

import (
	"github.com/gin-gonic/gin"
	"gobase/controllers/handlers"
)

func HandleUserRoutes(router *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler()
	router.GET("", userHandler.List)
}
