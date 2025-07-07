package dao

import (
	"auth/biz/db/mysql"
	"auth/biz/model/po"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type DomainDao struct {
	conn *mysql.DbConn
}

func NewDomainDao(tx ...*gorm.DB) *DomainDao {
	return &DomainDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (d *DomainDao) Create(ctx context.Context, po *po.Domain) error {
	return d.conn.WithContext(ctx).Create(po).Error
}

// update domain by domain_id
func (d *DomainDao) Update(ctx context.Context, po *po.Domain) error {
	return d.conn.WithContext(ctx).
		Omit("domain_id").
		Where("domain_id", po.DomainID).
		Updates(po).Error
}

func (d *DomainDao) Delete(ctx context.Context, domainID string) error {
	return d.conn.WithContext(ctx).
		Where("domain_id", domainID).
		Delete(&po.Domain{}).Error
}

func (d *DomainDao) QueryDetail(ctx context.Context, domainID string) (*po.Domain, error) {
	var result *po.Domain

	err := d.conn.WithContext(ctx).
		Where("domain_id", domainID).
		Take(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		hlog.CtxErrorf(ctx, "query detail err: %v", err)
		return nil, err
	}
	return result, nil
}

func (d *DomainDao) Count(ctx context.Context) (int64, error) {
	var count int64
	err := d.conn.WithContext(ctx).
		Model(&po.Domain{}).
		Count(&count).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "count err: %v", err)
		return 0, err
	}
	return count, nil
}

func (d *DomainDao) QueryList(ctx context.Context, limit, offset int) ([]po.Domain, error) {
	var result []po.Domain
	err := d.conn.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&result).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query list err: %v", err)
		return nil, err
	}
	return result, nil
}
