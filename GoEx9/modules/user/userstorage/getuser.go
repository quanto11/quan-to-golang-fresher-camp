package userstorage

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

func (s *sqlStore) GetDataById(ctx context.Context,
	id int,
) (usermodel.User, error) {
	var result usermodel.User

	db := s.db

	if err := db.Where("id = ?", id).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
