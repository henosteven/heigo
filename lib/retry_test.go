package lib

import (
	"testing"
	"math/rand"
	"time"
	"errors"
	"fmt"
)

func TestTryDo(t *testing.T) {

	a, b := 1, 2
	c := "hello"
	TryDo(1, func() error {
		return customFunc(a, b, c)
	}, time.Millisecond)

	caseList := []struct{
		attemptCount int
		fn Func
		err error
	} {
		{
			11,
			randResult,
			errors.New(OutMaxRetCountError),
		},
		{
			0,
			randResult,
			errors.New(OutMaxRetCountError),
		},
	}

	for _, val := range caseList {
		res := TryDo(val.attemptCount, func() error {
			return val.fn()
		}, time.Millisecond)

		if res.Error() != val.err.Error() {
			t.Errorf("expect:%v, get:%v", val.err, res)
		}
	}
}

func randResult() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)
	fmt.Printf("rand num: %v", num)
	if num < 80 {
		return errors.New("err")
	}
	return nil
}

func customFunc(a, b int, c string) error{
	fmt.Println(a, b, c)
	return nil
}