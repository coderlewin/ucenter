package ioc

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/coderlewin/ucenter/internal/web"
	"github.com/coderlewin/ucenter/internal/web/middleware"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"github.com/spf13/viper"
	"time"
)

func InitWebServer(mws []app.HandlerFunc, userHdl *web.UserHandler) *server.Hertz {
	engine := server.Default(
		server.WithHostPorts(viper.GetString("server.port")),
	)

	engine.Use(mws...)

	g := engine.Group(viper.GetString("server.prefix"))
	userHdl.ConfigRoutes(g)

	return engine
}

func CommonMiddlewares() []app.HandlerFunc {
	return []app.HandlerFunc{
		sessionHandlerFunc(),
		accessLog(),
		middleware.NewCheckSessionAuthMiddlewareBuilder().Build(),
	}
}

func sessionHandlerFunc() app.HandlerFunc {
	store := cookie.NewStore(
		[]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"),
		[]byte("o6jdlo2cb9f9pb6h46fjmllw481ldebj"),
	)

	// cookie 名字是 ssid
	return sessions.New("ssid", store)
}

func accessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Milliseconds()
		hlog.CtxTracef(c, "status=%d cost=%dms method=%s full_path=%s client_ip=%s host=%s",
			ctx.Response.StatusCode(), latency,
			ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
	}
}
