package common

import (
	"time"
	"strconv"
)

type HeiTrace struct {
	TraceID string
	CTraceID string
}

func GenTrace() (trace HeiTrace){
	trace = HeiTrace{
		CTraceID: GenTraceID(),
	}
	return
}

func GenTraceWithTraceID(traceID string ) (trace HeiTrace){
	trace = HeiTrace{
		TraceID: traceID,
	}
	return
}

func GenTraceID() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}