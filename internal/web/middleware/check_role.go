package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/internal/web/vo"
	"github.com/coderlewin/ucenter/pkg/core"
	"github.com/coderlewin/ucenter/pkg/errno"
	"net/http"
)

type CheckRoleMiddlewareBuilder struct {
	role int32
}

func NewCheckRoleMiddlewareBuilder(role int32) *CheckRoleMiddlewareBuilder {
	return &CheckRoleMiddlewareBuilder{role: role}
}

func (m *CheckRoleMiddlewareBuilder) Build() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		value, exists := c.Get(constants.LoginUser)
		loginUser := value.(*vo.UserVO)
		if !exists {
			core.SendResponse(c, errno.ErrUnauthorization, nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if loginUser.UserRole != m.role {
			core.SendResponse(c, errno.ErrForbidden.SetDescription("权限不足"), nil)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next(ctx)
	}
}
