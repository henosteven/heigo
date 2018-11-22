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

	params := r.URL.Query()
	userID := params.Get("userId")
	if userID == "" {
		ResponseFailed(w, USERID_EMPTY, struct{}{})
		return
	}

	id, err := strconv.Atoi(userID)
	if err !=nil {
		ResponseFailed(w, USERID_NOT_DIGIT, struct{}{})
		return
	}

	userName, err := model.GetUserNameByID(id)
	if err != nil {
		resp = err.Error()
	}  else {
		resp = userName
	}
	ResponseSuccess(w, resp)
}
