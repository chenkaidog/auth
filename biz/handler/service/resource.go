package service

import (
	"auth/biz/dao"
	"auth/biz/model/errs"
	"auth/biz/model/po"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

type ResourceCreateParam struct {
	DomainID    string
	Status      string
	Name        string
	Description string
}

type ResourceUpdateParam struct {
	ResourceID  string
	Status      string
	Name        string
	Description string
}

type ResourceDetail struct {
	ResourceID  string
	DomainID    string
	Status      string
	Name        string
	Description string
}

type ResourceQueryParam struct {
	Page, Size int
	DomainID   string
	Status     string
	Name       string
}

func CreateResource(ctx context.Context, param *ResourceCreateParam) errs.Error {
	if err := dao.NewResourceDao().Create(ctx, &po.Resource{
		ResourceID:  uuid.NewString(),
		DomainID:    param.DomainID,
		Status:      param.Status,
		Name:        param.Name,
		Description: param.Description,
	}); err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "resource already exist")
			return errs.ResourceNameDuplicatedErr
		}
		return errs.ServerError
	}

	return nil
}

func UpdateResource(ctx context.Context, param *ResourceUpdateParam) errs.Error {
	if err := dao.NewResourceDao().Update(ctx, &po.Resource{
		ResourceID:  param.ResourceID,
		Status:      param.Status,
		Name:        param.Name,
		Description: param.Description,
	}); err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "resource already exist")
			return errs.ResourceNameDuplicatedErr
		}
		return errs.ServerError
	}
	return nil
}

func DeleteResource(ctx context.Context, resourceId string) errs.Error {
	if err := dao.NewResourceDao().Delete(ctx, resourceId); err != nil {
		return errs.ServerError
	}
	return nil
}

func QueryResourceDetail(ctx context.Context, resourceId string) (*ResourceDetail, errs.Error) {
	resourceDetail, err := dao.NewResourceDao().QueryDetail(ctx, resourceId)
	if err != nil {
		return nil, errs.ServerError
	}
	if resourceDetail == nil {
		return nil, errs.ResourceNotExistErr
	}
	return &ResourceDetail{
		ResourceID:  resourceDetail.ResourceID,
		DomainID:    resourceDetail.DomainID,
		Status:      resourceDetail.Status,
		Name:        resourceDetail.Name,
		Description: resourceDetail.Description,
	}, nil
}

func QueryResourceList(ctx context.Context, param *ResourceQueryParam) ([]*ResourceDetail, int64, errs.Error) {
	total, err := dao.NewResourceDao().Count(ctx, param.DomainID, &dao.ResourceCond{
		Name:   param.Name,
		Status: param.Status,
	})
	if err != nil {
		return nil, 0, errs.ServerError
	}
	if total <= 0 {
		return nil, 0, nil
	}

	limit, offset := param.Size, (param.Page-1)*param.Size
	resources, err := dao.NewResourceDao().QueryList(ctx, limit, offset, param.DomainID, &dao.ResourceCond{
		Name:   param.Name,
		Status: param.Status,
	})
	if err != nil {
		return nil, 0, errs.ServerError
	}
	var result []*ResourceDetail
	for _, resource := range resources {
		result = append(result, &ResourceDetail{
			ResourceID:  resource.ResourceID,
			DomainID:    resource.DomainID,
			Status:      resource.Status,
			Name:        resource.Name,
			Description: resource.Description,
		})
	}
	return result, total, nil
}
