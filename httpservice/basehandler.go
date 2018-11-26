package httpservice

import (
	"encoding/json"
	"github.com/henosteven/heigo/common"
	"net/http"
	"runtime"
)

const (
	SUCCESS      = 0
	ERROR_COMMON = 1
)

const (
	SUCCESS_DESC      = "success"
	ERROR_COMMON_DESC = "something wrong happend"
)

type ResponseData struct {
	Code    int
	Message string
	Data    interface{}
}

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"errno\": 500}"))
				stack := make([]byte, 4<<10)
				runtime.Stack(stack, true)
				common.LogFatal(GetTraceInfoFromRequest(r), string(stack))
			}
		}()
		common.LogTrace(GetTraceInfoFromRequest(r), "com_request_in")
		if !common.LimitAllow(r.URL.Path) {
			panic("limit-qps")
		}
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

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	responseData := ResponseData{
		Code:    SUCCESS,
		Message: SUCCESS_DESC,
		Data:    data,
	}
	resp, _ := json.Marshal(responseData)
	w.Write(resp)
}

func ResponseFailed(w http.ResponseWriter, message string, data interface{}) {
	responseData := ResponseData{
		Code:    ERROR_COMMON,
		Message: message,
		Data:    data,
	}
	resp, _ := json.Marshal(responseData)
	w.Write(resp)
}
