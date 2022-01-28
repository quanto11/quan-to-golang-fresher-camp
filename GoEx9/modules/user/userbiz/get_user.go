package userbiz

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

type UserStore interface {
	GetDataById(ctx context.Context,
		id int,
	) (usermodel.User, error)
}

type userBiz struct {
	store UserStore
}

func NewUserBiz(store UserStore) *userBiz {
	return &userBiz{store: store}
}

func (biz *userBiz) GetUser(
	ctx context.Context,
	id int,
) (usermodel.User, error) {
	result, err := biz.store.GetDataById(ctx, id)

	return result, err
}
