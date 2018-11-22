package httpservice

import (
	"net/http"
	"github.com/henosteven/heigo/common"
)

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"errno\": 500}"))
			}
		}()
		common.LogTrace(GetTraceInfoFromRequest(r), "com_request_in")
		fn(w, r)
	}
}

func GetTraceInfoFromRequest(r *http.Request) common.HeiTrace {
	traceID := r.Header.Get("traceid")
	if traceID == "" {
		return common.GenTrace()
	} else {
		return common.GenTraceWithTraceID(traceID)
	}
}