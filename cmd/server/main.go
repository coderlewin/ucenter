package main

import (
	"encoding/gob"
	"github.com/coderlewin/ucenter/internal/web/vo"
	"github.com/coderlewin/ucenter/pkg/cfg"
)

func main() {
	// 加载配置
	if err := cfg.Load(); err != nil {
		panic(err)
	}
	gob.Register(&vo.UserVO{})

	app := InitApp()
	app.web.Spin()
}
