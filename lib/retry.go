package lib

import (
	"errors"
	"time"
)

const (
	MinRetryCount = 1
	MaxRetryCount = 10
	OutMaxRetCountError = "Out Of MaxLimit"
)
type Func func() error

type RetryStopErr struct {
	error
}

func TryDo(attempt int, fn Func, sleep time.Duration) (err error) {
	if attempt > MaxRetryCount || attempt < MinRetryCount {
		return errors.New(OutMaxRetCountError)
	}

	for i := 0; i <= attempt; i++ {

		err = fn()
		if err == nil {
			break
		}

		if _, ok := err.(RetryStopErr); ok {
			break
		}

		time.Sleep(sleep)
	}

	return err
}