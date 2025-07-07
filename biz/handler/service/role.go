package service

import (
	"auth/biz/dao"
	"auth/biz/model/errs"
	"auth/biz/model/po"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

type RoleCreateParam struct {
	DomainID    string
	ParentID    string
	Status      string
	Name        string
	Description string
}

type RoleUpdateParam struct {
	RoleID      string
	Status      string
	Name        string
	Description string
}

type RoleDetail struct {
	RoleID      string
	DomainID    string
	ParentID    string
	Status      string
	Name        string
	Description string
}

type RoleQueryParam struct {
	Page, Size int
	DomainID   string
	Status     string
	Name       string
}

func CreateRole(ctx context.Context, param *RoleCreateParam) errs.Error {
	parentRole, err := dao.NewRoleDao().QueryDetail(ctx, param.ParentID)
	if err != nil {
		return errs.ServerError
	}
	if parentRole == nil {
		hlog.CtxInfof(ctx, "parent role not exist")
		return errs.ParentRoleNotExistErr
	}
	if parentRole.DomainID != param.DomainID {
		hlog.CtxInfof(ctx, "parent role not in the same domain")
		return errs.ParentRoleNotInSameDomainErr
	}

	if err := dao.NewRoleDao().Create(ctx, &po.Role{
		RoleID:       uuid.NewString(),
		DomainID:     param.DomainID,
		ParentRoleID: param.ParentID,
		Status:       param.Status,
		Name:         param.Name,
		Description:  param.Description,
	}); err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "role already exist")
			return errs.RoleNameDuplicatedErr
		}

		return errs.ServerError
	}

	return nil
}

func UpdateRole(ctx context.Context, param *RoleUpdateParam) errs.Error {
	err := dao.NewRoleDao().Update(ctx, &po.Role{
		RoleID:      param.RoleID,
		Status:      param.Status,
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "role already exist")
			return errs.RoleNameDuplicatedErr
		}

		return errs.ServerError
	}

	return nil
}

func DeleteRole(ctx context.Context, roleId string) errs.Error {
	err := dao.NewRoleDao().Delete(ctx, roleId)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func QueryRoleDetail(ctx context.Context, roleId string) (*RoleDetail, errs.Error) {
	rolePO, err := dao.NewRoleDao().QueryDetail(ctx, roleId)
	if err != nil {
		return nil, errs.ServerError
	}
	if rolePO == nil {
		hlog.CtxInfof(ctx, "role not exist")
		return nil, errs.RoleNotExistErr
	}
	return &RoleDetail{
		RoleID:      rolePO.RoleID,
		DomainID:    rolePO.DomainID,
		ParentID:    rolePO.ParentRoleID,
		Status:      rolePO.Status,
		Name:        rolePO.Name,
		Description: rolePO.Description,
	}, nil
}

func QueryRoleList(ctx context.Context, param *RoleQueryParam) ([]*RoleDetail, int, errs.Error) {
	cnt, err := dao.NewRoleDao().Count(ctx, param.DomainID, &dao.RoleCond{
		Name:   param.Name,
		Status: param.Status,
	})
	if err != nil {
		return nil, 0, errs.ServerError
	}
	if cnt <= 0 {
		return nil, 0, nil
	}

	limit, offset := param.Size, (param.Page-1)*param.Size
	roleList, err := dao.NewRoleDao().QueryList(ctx, limit, offset, param.DomainID, &dao.RoleCond{
		Name:   param.Name,
		Status: param.Status,
	})
	if err != nil {
		return nil, 0, errs.ServerError
	}
	var roleDetailList []*RoleDetail
	for _, role := range roleList {
		roleDetailList = append(roleDetailList, &RoleDetail{
			RoleID:      role.RoleID,
			DomainID:    role.DomainID,
			ParentID:    role.ParentRoleID,
			Status:      role.Status,
			Name:        role.Name,
			Description: role.Description,
		})
	}
	return roleDetailList, int(cnt), nil
}
