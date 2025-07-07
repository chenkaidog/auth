package dao

import (
	"auth/biz/db/mysql"
	"auth/biz/model/po"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type AccountDao struct {
	conn *mysql.DbConn
}

func NewAccountDao(tx ...*gorm.DB) *AccountDao {
	return &AccountDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (d *AccountDao) Create(ctx context.Context, po *po.Account) error {
	return d.conn.WithContext(ctx).Create(po).Error
}

func (d *AccountDao) QueryByUsername(ctx context.Context, username string) (*po.Account, error) {
	var result *po.Account

	err := d.conn.WithContext(ctx).
		Where("username", username).
		Take(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query by username errs: %v", err)
		return nil, err
	}

	return result, nil
}
