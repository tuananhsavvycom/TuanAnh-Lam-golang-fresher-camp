package usersbiz

import (
	"Project/common"
	"Project/modules/user/usermodel"
	"context"
)

type GetUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return data, err
}
