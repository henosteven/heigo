package common

import (
	"time"
	"fmt"
)

type Limit struct {
	CurrentCount int
	CurrentTime int64
	LimitCount int
	LimitInterval int64
}

var LimitConfig map[string]*Limit

func InitLimitConfig() {
	LimitConfig = map[string]*Limit {
		"user" : {0, 0, 10, 1},
	}
}

func LimitAllow(api string) (allow bool) {
	allow = true
	limitConfig, ok := LimitConfig[api]
	if !ok {
		return
	}
	cur := time.Now().Unix()
	if cur > (limitConfig.CurrentTime + limitConfig.LimitInterval) {
		limitConfig.CurrentTime = cur
		limitConfig.CurrentCount = 1
		return
	}
	fmt.Println(limitConfig.CurrentCount)
	limitConfig.CurrentCount++
	if limitConfig.CurrentCount > limitConfig.LimitCount {
		allow = false
	}
	return
}