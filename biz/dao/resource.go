package dao

import (
	"auth/biz/db/mysql"
	"auth/biz/model/po"
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type ResourceDao struct {
	conn *mysql.DbConn
}

func NewResourceDao(tx ...*gorm.DB) *ResourceDao {
	return &ResourceDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (r *ResourceDao) Create(ctx context.Context, po *po.Resource) error {
	return r.conn.WithContext(ctx).Create(po).Error
}

// update resource by resource_id
func (r *ResourceDao) Update(ctx context.Context, po *po.Resource) error {
	return r.conn.WithContext(ctx).
		Omit("resource_id").
		Where("resource_id", po.ResourceID).
		Updates(po).Error
}

func (r *ResourceDao) Delete(ctx context.Context, resourceID string) error {
	return r.conn.WithContext(ctx).
		Where("resource_id", resourceID).
		Delete(&po.Resource{}).Error
}

func (r *ResourceDao) QueryDetail(ctx context.Context, resourceID string) (*po.Resource, error) {
	var result *po.Resource
	err := r.conn.WithContext(ctx).
		Where("resource_id", resourceID).
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

type ResourceCond struct {
	Name, Status string
}

func (r *ResourceDao) Count(ctx context.Context, domainId string, cond *ResourceCond) (int64, error) {
	var count int64

	sql := r.conn.WithContext(ctx).
		Model(&po.Resource{}).
		Where("domain_id", domainId)
	if cond != nil {
		if cond.Name != "" {
			sql = sql.Where("name Like ?", fmt.Sprintf("%%%s%%", cond.Name))
		}
		if cond.Status != "" {
			sql = sql.Where("status", cond.Status)
		}
	}
	err := sql.Count(&count).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "count err: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *ResourceDao) QueryList(ctx context.Context, limit, offset int, domainId string, cond *ResourceCond) ([]*po.Resource, error) {
	var result []*po.Resource
	sql := r.conn.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("name").
		Where("domain_id", domainId)
	if cond != nil {
		if cond.Name != "" {
			sql = sql.Where("name Like ?", fmt.Sprintf("%%%s%%", cond.Name))
		}
		if cond.Status != "" {
			sql = sql.Where("status", cond.Status)
		}
	}
	if err := sql.Find(&result).Error; err != nil {
		hlog.CtxErrorf(ctx, "query list err: %v", err)
		return nil, err
	}
	return result, nil
}
