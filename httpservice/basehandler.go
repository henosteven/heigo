package httpservice

import (
	"net/http"

)

func SafeHandler(fn http.HandlerFunc)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"errno\": 500}"))
			}
		}()
		fn(w, r)
	}
}