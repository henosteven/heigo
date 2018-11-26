package common

import (
	"fmt"
	"strconv"
	"time"
)

type HeiTrace struct {
	TraceID  string
	CTraceID string
}

func (this HeiTrace) GetTraceString() string {
	return fmt.Sprintf("trace:%s|ctraceid:%s", this.TraceID, this.CTraceID)
}

func GenTrace() (trace HeiTrace) {
	trace = HeiTrace{
		CTraceID: GenTraceID(),
	}
	return
}

func GenTraceWithTraceID(traceID string) (trace HeiTrace) {
	trace = HeiTrace{
		TraceID: traceID,
	}
	return
}

func GenTraceID() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}
