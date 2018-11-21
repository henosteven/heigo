package config

import (
	"time"
	"github.com/BurntSushi/toml"
	"fmt"
	"encoding/json"
)

const HOST = "127.0.0.1"

type RedisConfig struct {
	MaxIdle int	`json:"maxidle",toml:"maxidle"`
	IdleTimeout time.Duration	`json:"idle_timeout",toml:"idle_timeout"`
	Host string	`json:"host",toml:"host"`
	Port string	`json:"port",toml:"port"`
	User string	`json:"user",toml:"user"`
	Password string	`json:"password",toml:"password"`
}

type WebConfig struct {
	Port string `json:"port",toml:"port"`
}

type ThriftConfig struct {
	Port string	`json:"port",toml:"port"`
}

type MysqlConfig struct {
	Host string	`json:"host",toml:"host"`
	Port string	`json:"port",toml:"port"`
	User string	`json:"user",toml:"user"`
	Password string	`json:"password",toml:"password"`
	Database string	`json:"database",toml:"database"`
	Protocol string	`json:"protocol",toml:"protocol"`
}

type Config struct {
	Host string `json:"host",toml:"host"`
	WebConf WebConfig	`json:"web_conf",toml:"webconf"`
	ThriftConf ThriftConfig	`json:"thrift_conf",toml:"thriftconf`
	RedisConf  RedisConfig	`json:"redis_conf",toml:"redisconf"`
	MysqlConf  MysqlConfig	`json:"mysql_conf",toml:"mysqlconf"`
}

var GlobalConfig Config

func InitConfig(path string) {
	if _, err := toml.DecodeFile(path, &GlobalConfig); err != nil {
		panic(err)
	}
}

func ShowConfig() {
	val, err := json.Marshal(GlobalConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(val))
}