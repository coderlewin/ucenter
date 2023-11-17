package ioc

import (
	"fmt"
	"github.com/coderlewin/ucenter/pkg/db"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	mysql, err := db.NewMySQL(&db.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	})

	if err != nil {
		panic(fmt.Errorf("初始化配置失败, 原因 %w", err))
	}

	return mysql
}
