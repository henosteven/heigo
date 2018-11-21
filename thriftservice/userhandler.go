package thriftservice

import (
	"github.com/henosteven/heigo/heiThrift"
	"context"
)

type UserHandlerImpl struct {

}

func (userhandler UserHandlerImpl) GetUser(ctx context.Context, userId int64) (r *heiThrift.User, err error) {
	user := &heiThrift.User {
		UserId: 1,
		UserName:"jinjing",
	}
	return user, nil
}

func (userhandler UserHandlerImpl) AddUser(ctx context.Context, username string) (r int64, err error) {
	return 1, nil
}