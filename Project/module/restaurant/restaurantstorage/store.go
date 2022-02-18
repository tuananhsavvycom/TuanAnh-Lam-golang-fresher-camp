package restaurantstorage

import (
	"context"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	//TODO implement me
	panic("implement me")
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
