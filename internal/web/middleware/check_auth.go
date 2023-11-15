package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/pkg/core"
	"github.com/ecodeclub/ekit/set"
	"net/http"
)

type CheckSessionAuthMiddlewareBuilder struct {
	publicPaths set.Set[string]
}

func NewCheckSessionAuthMiddlewareBuilder() *CheckSessionAuthMiddlewareBuilder {
	s := set.NewMapSet[string](3)
	s.Add("/api/user/register")
	s.Add("/api/user/login")
	return &CheckSessionAuthMiddlewareBuilder{
		publicPaths: s,
	}
}

func (m *CheckSessionAuthMiddlewareBuilder) Build() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 不需要校验用户认证
		if m.publicPaths.Exist(string(ctx.Request.URI().Path())) {
			return
		}

		loginUser, err := core.GetUserLoginState(ctx)
		if err != nil {
			core.SendResponse(ctx, err, nil)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(constants.LoginUser, loginUser)

		ctx.Next(c)
	}
}
