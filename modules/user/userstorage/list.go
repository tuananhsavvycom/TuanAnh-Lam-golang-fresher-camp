package userstorage

import (
	"Project/common"
	"Project/modules/user/usermodel"
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

	if v := filter; v != nil {
		if v.User_id > 0 {
			db = db.Where("user_id = ?", v.User_id)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
