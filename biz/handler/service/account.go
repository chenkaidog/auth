package service

import (
	"auth/biz/dao"
	"auth/biz/model/consts"
	"auth/biz/model/errs"
	"auth/biz/util/encode"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type LoginParam struct {
	Username, Password string
}

type LoginResult struct {
	UserID string
	Name   string
}

func Login(ctx context.Context, param LoginParam) (LoginResult, errs.Error) {
	accPO, err := dao.NewAccountDao().QueryByUsername(ctx, param.Username)
	if err != nil {
		return LoginResult{}, errs.ServerError
	}
	if accPO == nil {
		hlog.CtxInfof(ctx, "account not exist")
		return LoginResult{}, errs.AccountNotExist
	}

	if encode.EncodePassword(accPO.Salt, param.Password) != accPO.Password {
		hlog.CtxInfof(ctx, "password incorrect")
		return LoginResult{}, errs.PasswordIncorrect
	}
	if consts.AccountStatusValid != accPO.Status {
		hlog.CtxInfof(ctx, "account status invalid: %s", accPO.Status)
		return LoginResult{}, errs.AccountNotExist
	}

	userPO, err := dao.NewUserDao().QueryByAccountId(ctx, accPO.AccountID)
	if err != nil || userPO == nil {
		return LoginResult{}, errs.ServerError
	}
	return LoginResult{
		UserID: userPO.UserID,
		Name:   userPO.Name,
	}, nil
}
