package common

import (
	"testing"
)

func TestMain(m *testing.M) {
	InitLimitConfig()
	m.Run()
}

func TestLimitAllow(t *testing.T) {
	caseList := []struct {
		Api   string
		Allow bool
	}{
		{"user", true},
	}

	for i := 0; i < 11; i++ {
		for _, val := range caseList {
			result := LimitAllow(val.Api)
			if result != val.Allow {
				t.Fatalf("failed api:%s expect: %v, result:%v", val.Api, val.Allow, result)
			}
		}
	}
}
