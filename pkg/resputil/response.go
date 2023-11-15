package resputil

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type BaseResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Data        any    `json:"data"`
}

func Ok(c *app.RequestContext, data any) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:        0,
		Message:     "ok",
		Description: "",
		Data:        data,
	})
}

func Fail(c *app.RequestContext, code int, message, description string) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:        code,
		Message:     message,
		Description: description,
		Data:        nil,
	})
}
