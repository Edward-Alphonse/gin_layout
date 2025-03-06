package config

import (
	"gin_layout/internal/libs/config"
	"gin_layout/internal/mysql"
	"gin_layout/internal/redis"
)

type App struct {
	Port string
}

type Configuration struct {
	App   App
	Mysql *mysql.Config
	Redis *redis.Config
}

var Cfg *Configuration

func Init(path string) *Configuration {
	Cfg = config.Init[Configuration](path)
	return Cfg
}
