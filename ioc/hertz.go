package ioc

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/coderlewin/ucenter/internal/web"
	"github.com/coderlewin/ucenter/internal/web/middleware"
	"github.com/coderlewin/ucenter/pkg/cfg"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func InitWebServer(mws []app.HandlerFunc, userHdl *web.UserHandler) *server.Hertz {
	engine := server.Default(
		server.WithHostPorts(cfg.MustGet[string]("server.port")),
	)

	engine.Use(mws...)

	g := engine.Group(cfg.MustGet[string]("server.prefix"))
	userHdl.ConfigRoutes(g)

	return engine
}

func CommonMiddlewares() []app.HandlerFunc {
	return []app.HandlerFunc{
		sessionHandlerFunc(),
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
