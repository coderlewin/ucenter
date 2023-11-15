package core

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/coderlewin/ucenter/pkg/resputil"
)

func SendResponse(c *app.RequestContext, err error, data any) {
	if err != nil {
		code, msg, desc := errno.Decode(err)
		resputil.Fail(c, code, msg, desc)
		return
	}
	resputil.Ok(c, data)
}
