package db

import "fmt"

type Config struct {
	UserName     string
	Password     string
	Address      string
	DbName       string
	Charset      string
	TimeOut      string //连接超时，10秒
	MaxOpenConns int    // 最大连接数
	MaxIdleConns int    // 最大空闲连接数
}

func (d *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&timeout=%s", d.UserName, d.Password, d.Address, d.DbName, d.Charset, d.TimeOut)
}
