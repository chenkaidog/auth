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

type RoleDao struct {
	conn *mysql.DbConn
}

func NewRoleDao(tx ...*gorm.DB) *RoleDao {
	return &RoleDao{
		conn: mysql.NewDbConn(tx...),
	}
}

func (r *RoleDao) Create(ctx context.Context, po *po.Role) error {
	return r.conn.WithContext(ctx).Create(po).Error
}

// update role by role_id
func (r *RoleDao) Update(ctx context.Context, po *po.Role) error {
	return r.conn.WithContext(ctx).
		Omit("role_id").
		Where("role_id", po.RoleID).
		Updates(po).Error
}

func (r *RoleDao) Delete(ctx context.Context, roleID string) error {
	return r.conn.WithContext(ctx).
		Where("role_id", roleID).
		Delete(&po.Role{}).Error
}

func (r *RoleDao) QueryDetail(ctx context.Context, roleID string) (*po.Role, error) {
	var result *po.Role
	err := r.conn.WithContext(ctx).
		Where("role_id", roleID).
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

type RoleCond struct {
	Name, Status string
}

func (r *RoleDao) Count(ctx context.Context, domainId string, cond *RoleCond) (int64, error) {
	var count int64

	sql := r.conn.WithContext(ctx).
		Model(&po.Role{}).
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

func (r *RoleDao) QueryList(ctx context.Context, limit, offset int, domainId string, cond *RoleCond) ([]*po.Role, error) {
	var result []*po.Role

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

	err := sql.Find(&result).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query list err: %v", err)
		return nil, err
	}
	return result, nil
}
