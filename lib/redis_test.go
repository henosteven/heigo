package lib

import (
	"time"
	"testing"
	"github.com/henosteven/heigo/config"
)

var conf = config.RedisConfig{
	MaxIdle: 3,
	IdleTimeout: time.Second * 1,
	Host: "127.0.0.1",
	Port: "6379",
}

func TestMain(m *testing.M) {
	InitRedis(conf)
	m.Run()
}

func TestSet(t *testing.T) {
	var demo = map[string]string {
		"name": "test",
	}

	for key, val := range(demo) {
		err := Set(key, val)
		if err != nil {
			t.Error("case failed: set", err)
		}
	}
}

func TestGet(t *testing.T) {
	var demo = map[string]string {
		"name": "test",
	}

	for key, val := range (demo) {
		cacheval, err := Get(key)

		if err != nil {
			t.Error("case failed: get")
		}

		if cacheval != val {
			t.Error("case failed: get content failed")
		}
	}
}
