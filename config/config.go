package config

import "time"

const HOST = "127.0.0.1"
const WEB_PORT = "3002"
const THRIFT_PORT = "3001"


type RedisConfig struct {
	MaxIdle int
	IdleTimeout time.Duration
	Host string
	Port string
	User string
	Password string
}