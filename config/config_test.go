package config

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig("./conf.toml")
	ShowConfig()
}
