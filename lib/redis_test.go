package lib

import (
	"time"
	"testing"
	"github.com/henosteven/heigo/config"
	"os"
)

var conf = config.RedisConfig{
	MaxIdle: 3,
	IdleTimeout: time.Second * 1,
	Host: "127.0.0.1",
	Port: "6379",
}

func TestMain(m *testing.M) {
	InitRedis(conf)
	retCode := m.Run()
	os.Exit(retCode)
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
	var demo = []struct {
		key string
		expect string
	} {
		{"name", "test"},
		{"name1", ""},
	}

	for _, demoItem := range demo {
		cacheVal, err := Get(demoItem.key)

		if err != nil {
			t.Errorf("case failed err: %v", err)
		}

		if cacheVal != demoItem.expect {
			t.Errorf("case failed: get  key:%s, expect: %s, get:%s", demoItem.key, demoItem.expect, cacheVal)
		}
	}
}
