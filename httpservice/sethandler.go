package httpservice

import (
	"github.com/henosteven/heigo/model"
	"net/http"
)

func SetUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	userID, err := model.AddUser(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	ResponseSuccess(w, userID)
}
