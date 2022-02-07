package userstorage

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (usermodel.User, error) {
	var result usermodel.User

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
