package userstorage

import (
	"Project/modules/user/usermodel"
	"context"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	var result usermodel.User

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
