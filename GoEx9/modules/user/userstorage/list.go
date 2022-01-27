package userstorage

import (
	"GoEx8/common"
	"GoEx8/modules/user/usermodel"
	"context"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(usermodel.User{}.TableName()).Where(conditions)

	if v := filter; v.Status == 0 || v.Status == 1 {
		db = db.Where("status = ?", v.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
