package usersbiz

import (
	"Project/common"
	"Project/modules/user/usermodel"
	"context"
)

type ListUserStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	return result, err
}
