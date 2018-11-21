package httpservice

import (
	"github.com/henosteven/heigo/model"
	"net/http"
	"strconv"
)

const (
	USERID_EMPTY = "user id empty"
	USERID_NOT_DIGIT = "user id not digit"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userID := params.Get("userId")
	if userID == "" {
		w.Write([]byte(USERID_EMPTY))
	}

	id, err := strconv.Atoi(userID)
	if err !=nil {
		w.Write([]byte(USERID_NOT_DIGIT))
	}

	userName, err := model.GetUserNameByID(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}  else {
		w.Write([]byte(userName))
	}
}
