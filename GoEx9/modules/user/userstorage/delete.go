package userstorage

import (
	"GoEx8/modules/user/usermodel"
	"context"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}

	return nil
}
