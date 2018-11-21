package httpservice

import "net/http"

func SetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yes~~just~for~test"))
}