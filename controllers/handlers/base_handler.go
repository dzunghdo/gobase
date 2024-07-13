package handlers

import (
	"github.com/gin-gonic/gin"
	"gobase/controllers/middlewares"
)

type BaseHandler struct{}

func (h *BaseHandler) SetData(ctx *gin.Context, data interface{}) {
	ctx.Set(middlewares.CtxKeyData, data)
}

func (h *BaseHandler) SetError(ctx *gin.Context, err error) {
	ctx.Set(middlewares.CtxKeyError, err)
}

func (h *BaseHandler) SetStatus(ctx *gin.Context, status int) {
	ctx.Set(middlewares.CtxKeyStatus, status)
}
