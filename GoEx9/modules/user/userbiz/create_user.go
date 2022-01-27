package userbiz

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

type CreateUserStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUser(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}
