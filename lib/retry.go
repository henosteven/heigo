package lib

import (
	"errors"
	"fmt"
)

const (
	MinRetryCount = 1
	MaxRetryCount = 10
	OutMaxRetCountError = "Out Of MaxLimit"
)
type Func func() error

func TryDo(attempt int, fn Func) (err error) {
	if attempt > MaxRetryCount || attempt < MinRetryCount {
		return errors.New(OutMaxRetCountError)
	}

	for i := 0; i <= attempt; i++ {
		fmt.Printf("try:%v", i)
		err = fn()
		if err == nil {
			break
		}
	}

	return err
}