package main

import (
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	if err := initViperConfig(); err != nil {
		panic(err)
	}

	app := InitApp()
	app.web.Spin()
}

// 根据 viper 初始化配置
func initViperConfig() error {
	viper.SetConfigFile("configs/application.yml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
