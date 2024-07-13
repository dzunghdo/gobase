package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CtxKeyData   = "context_data"
	CtxKeyError  = "context_error"
	CtxKeyStatus = "context_status"
)

//type BaseResponseWriter struct {
//	gin.ResponseWriter
//	body *bytes.Buffer
//}

//func (w *BaseResponseWriter) Write(b []byte) (int, error) {
//	w.body.Write(b)
//	return w.ResponseWriter.Write(b)
//}
//
//func (w BaseResponseWriter) WriteString(s string) (int, error) {
//	w.body.WriteString(s)
//	return w.ResponseWriter.WriteString(s)
//}

func ResponseMiddleware(ctx *gin.Context) {
	//writer := &BaseResponseWriter{
	//	body:           bytes.NewBufferString(""),
	//	ResponseWriter: ctx.Writer,
	//}
	//ctx.Writer = writer
	ctx.Next()

	response := BaseResponse{}
	httpStatus := ctx.GetInt(CtxKeyStatus)

	if data, hasData := ctx.Get(CtxKeyData); hasData {
		response.Data = data
		if httpStatus == 0 {
			if ctx.Request.Method == http.MethodPost {
				httpStatus = http.StatusCreated
			} else {
				httpStatus = http.StatusOK
			}
		}
	}

	if err, hasError := ctx.Get(CtxKeyError); hasError {
		response.Errors = err
		if httpStatus == 0 {
			httpStatus = http.StatusInternalServerError
		}
	}

	ctx.JSON(httpStatus, response)
}

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message"`
}
