package service

import (
	"auth/biz/dao"
	"auth/biz/model/errs"
	"context"
)

type UserInfo struct {
	Name        string
	Email       string
	Phone       string
	Description string
}

func GetUserInfo(ctx context.Context, userID string) (UserInfo, errs.Error) {
	userPO, err := dao.NewUserDao().QueryByUserId(ctx, userID)
	if err != nil || userPO == nil {
		return UserInfo{}, errs.ServerError
	}
	return UserInfo{
		Name:        userPO.Name,
		Email:       userPO.Email,
		Phone:       userPO.Phone,
		Description: userPO.Description,
	}, nil
}
