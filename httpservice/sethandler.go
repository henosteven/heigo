package httpservice

import (
	"net/http"
	"jinjing.space/web/model"
	"strconv"
)

func SetUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	username := values.Get("username")
	userID, err := model.AddUser(username)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("success" + strconv.Itoa(userID)))
}