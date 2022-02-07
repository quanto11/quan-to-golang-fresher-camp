package userbiz

import (
	"GoEx8/modules/user/usermodel"
	"context"
	"errors"
)

type UpdateUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (usermodel.User, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *usermodel.UserUpdate,
	) error
}

type updateUserBiz struct {
	store UpdateUserStore
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate,
) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
