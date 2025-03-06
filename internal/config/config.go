package config

import (
	"gin_layout/internal/db"
	"gin_layout/internal/libs/config"
	"gin_layout/internal/redis"
)

type App struct {
	Port string
}

type Configuration struct {
	App   App
	Mysql *db.Config
	Redis *redis.Config
}

var Cfg *Configuration

func Init(path string) *Configuration {
	Cfg = config.Init[Configuration](path)
	return Cfg
}
