package service

import (
	"auth/biz/dao"
	"auth/biz/model/errs"
	"auth/biz/model/po"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/google/uuid"
)

type DomainCreateParam struct {
	Name        string
	Description string
}

type DomainUpdateParam struct {
	DomainID    string
	Name        string
	Description string
}

type DomainDeleteParam struct {
	DomainID string
}

type DomainDetailQueryParam struct {
	DomainID string
}

type DomainDetail struct {
	DomainID    string
	Name        string
	Description string
}

type DomainListQueryParam struct {
	Page, Size int
}

func CreateDomain(ctx context.Context, param *DomainCreateParam) errs.Error {
	err := dao.NewDomainDao().Create(ctx, &po.Domain{
		DomainID:    uuid.NewString(),
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "domain already exist")
			return errs.DomainNameDuplicatedErr
		}
		return errs.ServerError
	}

	return nil
}

func UpdateDomain(ctx context.Context, param *DomainUpdateParam) errs.Error {
	err := dao.NewDomainDao().Update(ctx, &po.Domain{
		DomainID:    param.DomainID,
		Name:        param.Name,
		Description: param.Description,
	})
	if err != nil {
		if errs.IsDuplicatedErr(err) {
			hlog.CtxInfof(ctx, "domain already exist")
			return errs.DomainNameDuplicatedErr
		}
		return errs.ServerError
	}

	return nil
}

func DeleteDomain(ctx context.Context, param *DomainDeleteParam) errs.Error {
	err := dao.NewDomainDao().Delete(ctx, param.DomainID)
	if err != nil {
		return errs.ServerError
	}
	return nil
}

func QueryDomainDetail(ctx context.Context, param *DomainDetailQueryParam) (*DomainDetail, errs.Error) {
	domainPO, err := dao.NewDomainDao().QueryDetail(ctx, param.DomainID)
	if err != nil {
		return nil, errs.ServerError
	}
	if domainPO == nil {
		hlog.CtxInfof(ctx, "domain not exist")
		return nil, errs.DomainNotExistErr
	}

	return &DomainDetail{
		DomainID:    domainPO.DomainID,
		Name:        domainPO.Name,
		Description: domainPO.Description,
	}, nil
}

func QueryDomainList(ctx context.Context, param *DomainListQueryParam) ([]*DomainDetail, int64, errs.Error) {
	cnt, err := dao.NewDomainDao().Count(ctx)
	if err != nil {
		return nil, 0, errs.ServerError
	}
	if cnt <= 0 {
		return nil, 0, nil
	}

	limit, offset := param.Size, (param.Page-1)*param.Size
	domainPOList, err := dao.NewDomainDao().QueryList(ctx, limit, offset)
	if err != nil {
		return nil, 0, errs.ServerError
	}
	var result []*DomainDetail
	for _, domainPO := range domainPOList {
		result = append(result, &DomainDetail{
			DomainID:    domainPO.DomainID,
			Name:        domainPO.Name,
			Description: domainPO.Description,
		})
	}
	return result, cnt, nil
}
