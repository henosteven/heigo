package httpservice

import (
	"github.com/henosteven/heigo/model"
	"net/http"
	"strconv"
)

func SetUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	userID, err := model.AddUser(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("success" + strconv.Itoa(userID)))
}