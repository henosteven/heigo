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
	var resp string

	response:
		w.Write([]byte(resp))
		return

	params := r.URL.Query()
	userID := params.Get("userId")
	if userID == "" {
		resp = USERID_EMPTY
		goto response
	}

	id, err := strconv.Atoi(userID)
	if err !=nil {
		resp = USERID_NOT_DIGIT
		goto response
	}

	userName, err := model.GetUserNameByID(id)
	if err != nil {
		resp = err.Error()
	}  else {
		resp = userName
	}
	goto response
}
