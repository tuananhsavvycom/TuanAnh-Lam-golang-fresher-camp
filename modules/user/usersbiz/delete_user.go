package usersbiz

import (
	"Project/modules/user/usermodel"
	"context"
	"errors"
)

type DeleteUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*usermodel.User, error)

	SoftDeleteData(
		ctx context.Context,
		user_id int,
	) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, user_id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": user_id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDeleteData(ctx, user_id); err != nil {
		return nil
	}

	return nil
}
