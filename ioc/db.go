package ioc

import (
  "fmt"
  "github.com/coderlewin/ucenter/pkg/cfg"
  "github.com/coderlewin/ucenter/pkg/db"
  "gorm.io/gorm"
  "time"
)

func InitDB() *gorm.DB {
  mysql, err := db.NewMySQL(&db.MySQLOptions{
    Host:                  cfg.MustGet[string]("db.host"),
    Username:              cfg.MustGet[string]("db.username"),
    Password:              cfg.MustGet[string]("db.password"),
    Database:              cfg.MustGet[string]("db.database"),
    MaxIdleConnections:    cfg.MustGet[int]("db.max-idle-connections"),
    MaxOpenConnections:    cfg.MustGet[int]("db.max-open-connections"),
    MaxConnectionLifeTime: time.Duration(cfg.MustGet[int]("db.max-connection-life-time")) * time.Second,
    LogLevel:              cfg.MustGet[int]("db.log-level"),
  })

  if err != nil {
    panic(fmt.Errorf("初始化配置失败, 原因 %w", err))
  }

  return mysql
}
