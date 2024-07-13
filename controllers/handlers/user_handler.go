package handlers

import (
	"github.com/gin-gonic/gin"
	"gobase/services/users"
)

type UserHandler struct {
	BaseHandler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (userHandler *UserHandler) List(ctx *gin.Context) {
	var err error
	defer func() {
		userHandler.SetError(ctx, err)
	}()

	users, err := users.NewUserUseCase().List()
	if err != nil {
		return
	}
	userHandler.SetData(ctx, users)
}
