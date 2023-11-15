package core

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/coderlewin/ucenter/internal/constants"
	"github.com/coderlewin/ucenter/internal/web/vo"
	"github.com/coderlewin/ucenter/pkg/errno"
	"github.com/hertz-contrib/sessions"
)

func SetUserLoginState(c *app.RequestContext, data any) error {
	session := sessions.Default(c)
	sessionTTL := 86400
	session.Set(constants.UserLoginState, data)
	// 设置过期时间
	session.Options(sessions.Options{MaxAge: sessionTTL})
	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserLoginState(c *app.RequestContext) error {
	session := sessions.Default(c)
	session.Delete(constants.UserLoginState)
	session.Options(sessions.Options{MaxAge: -1})
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetUserLoginState(c *app.RequestContext) (*vo.UserVO, error) {
	session := sessions.Default(c)
	obj := session.Get(constants.UserLoginState)
	if obj == nil {
		return nil, errno.ErrUnauthorization.SetDescription("未登录")
	}
	return obj.(*vo.UserVO), nil
}
