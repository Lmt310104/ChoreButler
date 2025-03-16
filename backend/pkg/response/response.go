package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
