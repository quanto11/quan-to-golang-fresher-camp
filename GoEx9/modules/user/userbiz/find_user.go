package userbiz

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

type UserStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
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
	condition map[string]interface{},
	id int,
	moreKeys ...string,
) (usermodel.User, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return result, err
}
