package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gobase/services/auth"
	"gobase/services/auth/dto"
)

type AuthHandler struct {
	BaseHandler
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var (
		err        error
		httpStatus int
		req        dto.LoginRequest
	)
	defer func() {
		ah.SetError(ctx, err)
		ah.SetStatus(ctx, httpStatus)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		httpStatus = http.StatusBadRequest
		return
	}

	authServ := auth.NewAuthService()
	token, err := authServ.Login(ctx, req)
	if err != nil {
		httpStatus = http.StatusUnauthorized
		return
	}
	ah.SetData(ctx, token)
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var (
		err        error
		httpStatus int
		req        dto.RegisterRequest
	)
	defer func() {
		ah.SetError(ctx, err)
		ah.SetStatus(ctx, httpStatus)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	authServ := auth.NewAuthService()
	user, err := authServ.Register(ctx, req)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	ah.SetData(ctx, user)
}
