//go:build wireinject

package main

import (
	"github.com/coderlewin/ucenter/internal/infrastructure/persistence/mysql"
	"github.com/coderlewin/ucenter/internal/repository"
	"github.com/coderlewin/ucenter/internal/service"
	"github.com/coderlewin/ucenter/internal/web"
	"github.com/coderlewin/ucenter/ioc"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		ioc.InitDB,
		ioc.InitRedis,

		// DAO 部分
		mysql.NewUserDao,
		// Cache 部分

		// repository 部分
		repository.NewUserRepository,

		// service 部分
		service.NewUserService,

		// handler 部分
		web.NewUserHandler,

		// hertz 的中间件
		ioc.CommonMiddlewares,

		// Web 服务器
		ioc.InitWebServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
