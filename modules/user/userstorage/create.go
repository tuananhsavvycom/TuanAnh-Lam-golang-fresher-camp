package userstorage

import (
	"Project/modules/user/usermodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
