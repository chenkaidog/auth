package dao

import (
	"auth/biz/db/mysql"
	"auth/biz/model/po"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type UserDao struct {
	conn *mysql.DbConn
}

func NewUserDao(tx ...*gorm.DB) *UserDao {
	return &UserDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (d *UserDao) Create(ctx context.Context, po *po.User) error {
	return d.conn.WithContext(ctx).Create(po).Error
}

func (d *UserDao) QueryByAccountId(ctx context.Context, accID string) (*po.User, error) {
	var result *po.User

	err := d.conn.WithContext(ctx).
		Where("account_id", accID).
		Take(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by account id errs: %v", err)
		return nil, err
	}

	return result, nil
}

func (d *UserDao) QueryByUserId(ctx context.Context, userID string) (*po.User, error) {
	var result *po.User

	err := d.conn.WithContext(ctx).
		Where("user_id", userID).
		Take(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by user id errs: %v", err)
		return nil, err
	}

	return result, nil
}
