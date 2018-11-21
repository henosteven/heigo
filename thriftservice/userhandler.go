package thriftservice

import (
	"github.com/henosteven/heigo/heiThrift"
	"github.com/henosteven/heigo/model"
	"context"
)

type UserHandlerImpl struct {

}

func (userhandler UserHandlerImpl) GetUser(ctx context.Context, userId int64) (r *heiThrift.User, err error) {
	user := &heiThrift.User {}

	userName, err := model.GetUserNameByID(int(userId))
	if err != nil {
		return user, err
	}

	user.UserName = userName
	user.UserId = userId
	return user, nil
}

func (userhandler UserHandlerImpl) AddUser(ctx context.Context, username string) (r int64, err error) {
	userID, err := model.AddUser(username)
	return int64(userID), err
}