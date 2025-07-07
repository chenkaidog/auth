package handler

import (
	"auth/biz/handler/service"
	"auth/biz/model/dto"
	"auth/biz/model/errs"
	"auth/biz/util/convertor"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// CreateDomain
//
//	@Summary		创建作用域
//	@Description	创建作用域
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.DomainCreateReq	true	"create domain request body"
//	@Success		200	{object}	dto.CommonResp{data=dto.DomainCreateResp}
//	@Router			/api/v1/domain/create [POST]
func CreateDomain(ctx context.Context, c *app.RequestContext) {
	var domainCreateReq dto.DomainCreateReq
	if stdErr := c.BindAndValidate(&domainCreateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	bizErr := service.CreateDomain(ctx, &service.DomainCreateParam{
		Name:        domainCreateReq.Name,
		Description: domainCreateReq.Description,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	dto.SuccessResp(c, &dto.DomainCreateResp{})
}

// UpdateDomain
//
//	@Summary		更新作用域
//	@Description	更新作用域
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			domain_id	query		string				true	"update domain request body"
//	@Param			req			body		dto.DomainUpdateReq	true	"update domain request body"
//	@Success		200			{object}	dto.CommonResp{data=dto.DomainUpdateResp}
//	@Router			/api/v1/domain/update [POST]
func UpdateDomain(ctx context.Context, c *app.RequestContext) {
	var domainUpdateReq dto.DomainUpdateReq
	if stdErr := c.BindAndValidate(&domainUpdateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	bizErr := service.UpdateDomain(ctx, &service.DomainUpdateParam{
		DomainID:    domainUpdateReq.DomainID,
		Name:        domainUpdateReq.Name,
		Description: domainUpdateReq.Description,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.DomainUpdateResp{})
}

// DeleteDomain
//
//	@Summary		删除作用域
//	@Description	删除作用域
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			domain_id	query		string	true	"delete domain request body"
//	@Success		200			{object}	dto.CommonResp{data=dto.DomainDeleteResp}
//	@Router			/api/v1/domain/delete/ [POST]
func DeleteDomain(ctx context.Context, c *app.RequestContext) {
	var domainDeleteReq dto.DomainDeleteReq
	if stdErr := c.BindAndValidate(&domainDeleteReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	bizErr := service.DeleteDomain(ctx, &service.DomainDeleteParam{
		DomainID: domainDeleteReq.DomainID,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.DomainDeleteResp{})
}

// QueryDomainDetail
//
//	@Summary		查询作用域详情
//	@Description	查询作用域详情
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.DomainDetailQueryReq	true	"query domain request body"
//	@Success		200	{object}	dto.CommonResp{data=dto.DomainDetailQueryResp}
//	@Router			/api/v1/domain/query [GET]
func QueryDomainDetail(ctx context.Context, c *app.RequestContext) {
	var domainDetailQueryReq dto.DomainDetailQueryReq
	if stdErr := c.BindAndValidate(&domainDetailQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}
	domainDetail, bizErr := service.QueryDomainDetail(ctx, &service.DomainDetailQueryParam{
		DomainID: domainDetailQueryReq.DomainID,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.DomainDetailQueryResp{
		Domain: dto.Domain{
			DomainID:    domainDetail.DomainID,
			Name:        domainDetail.Name,
			Description: domainDetail.Description,
		},
	})
}

// QueryDomainList
//
//	@Summary		查询作用域列表
//	@Description	查询作用域列表
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.DomainListQueryReq	true	"query domain request body"
//	@Success		200	{object}	dto.CommonResp{data=dto.DomainListQueryResp}
//	@Router			/api/v1/domain/list [GET]
func QueryDomainList(ctx context.Context, c *app.RequestContext) {
	var domainListQueryReq dto.DomainListQueryReq
	if stdErr := c.BindAndValidate(&domainListQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	domainList, cnt, bizErr := service.QueryDomainList(ctx, &service.DomainListQueryParam{
		Page: domainListQueryReq.Page,
		Size: domainListQueryReq.Size,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	domainListQueryResp := &dto.DomainListQueryResp{
		Page:  domainListQueryReq.Page,
		Size:  domainListQueryReq.Size,
		Total: int(cnt),
	}
	for _, domain := range domainList {
		domainListQueryResp.Domains = append(domainListQueryResp.Domains, &dto.Domain{
			DomainID:    domain.DomainID,
			Name:        domain.Name,
			Description: domain.Description,
		})
	}
	dto.SuccessResp(c, domainListQueryResp)
}

// CreateRole
//
//	@Summary		创建角色
//	@Description	创建角色
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			domain_id	query		string				true	"create role under domain"
//	@Param			role		body		dto.RoleCreateReq	true	"create role"
//	@Success		200			{object}	dto.CommonResp{data=dto.RoleCreateResp}
//	@Router			/api/v1/domain/create_role [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var roleCreateReq dto.RoleCreateReq
	if stdErr := c.BindAndValidate(&roleCreateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	bizErr := service.CreateRole(ctx, &service.RoleCreateParam{
		DomainID:    roleCreateReq.DomainID,
		ParentID:    roleCreateReq.ParentID,
		Status:      convertor.RoleStatus2DO[roleCreateReq.Status],
		Name:        roleCreateReq.Name,
		Description: roleCreateReq.Description,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.RoleCreateResp{})
}

// QueryRoleList
//
//	@Summary		查询角色列表
//	@Description	查询角色列表
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.RoleListQueryReq	true	"query role list"
//	@Success		200	{object}	dto.CommonResp{data=dto.RoleListQueryResp}
//	@Router			/api/v1/domain/query_role [GET]
func QueryRoleList(ctx context.Context, c *app.RequestContext) {
	var roleListQueryReq dto.RoleListQueryReq
	if stdErr := c.BindAndValidate(&roleListQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	results, total, bizErr := service.QueryRoleList(ctx, &service.RoleQueryParam{
		Page:     roleListQueryReq.Page,
		Size:     roleListQueryReq.Size,
		DomainID: roleListQueryReq.DomainID,
		Status:   convertor.RoleStatus2DO[roleListQueryReq.Status],
		Name:     roleListQueryReq.Name,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	roleListQueryResp := &dto.RoleListQueryResp{
		Page:  roleListQueryReq.Page,
		Size:  roleListQueryReq.Size,
		Total: int(total),
	}
	for _, role := range results {
		roleListQueryResp.Roles = append(roleListQueryResp.Roles, &dto.Role{
			RoleID:      role.RoleID,
			DomainID:    role.DomainID,
			ParentID:    role.ParentID,
			Name:        role.Name,
			Status:      convertor.RoleStatus2DTO[role.Status],
			Description: role.Description,
		})
	}
	dto.SuccessResp(c, roleListQueryResp)
}

// CreateResource
//
//	@Summary		创建资源
//	@Description	创建资源
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			domain_id	query		string					true	"create resource under domain"
//	@Param			resource	body		dto.ResourceCreateReq	true	"create resource request body"
//	@Success		200			{object}	dto.CommonResp{data=dto.ResourceCreateResp}
//	@Router			/api/v1/domain/create_resource [POST]
func CreateResource(ctx context.Context, c *app.RequestContext) {
	var resourceCreateReq dto.ResourceCreateReq
	if stdErr := c.BindAndValidate(&resourceCreateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	bizErr := service.CreateResource(ctx, &service.ResourceCreateParam{
		DomainID:    resourceCreateReq.DomainID,
		Status:      convertor.ResourceStatus2DO[resourceCreateReq.Status],
		Name:        resourceCreateReq.Name,
		Description: resourceCreateReq.Description,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.ResourceCreateResp{})
}

// QueryResourceList
//
//	@Summary		查询资源列表
//	@Description	查询资源列表
//	@Tags			domain
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.ResourceListQueryReq	true	"query resource list"
//	@Success		200	{object}	dto.CommonResp{data=dto.ResourceListQueryResp}
//	@Router			/api/v1/domain/query_resource [GET]
func QueryResourceList(ctx context.Context, c *app.RequestContext) {
	var resourceListQueryReq dto.ResourceListQueryReq
	if stdErr := c.BindAndValidate(&resourceListQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	results, total, bizErr := service.QueryResourceList(ctx, &service.ResourceQueryParam{
		Page:     resourceListQueryReq.Page,
		Size:     resourceListQueryReq.Size,
		DomainID: resourceListQueryReq.DomainID,
		Status:   convertor.ResourceStatus2DO[resourceListQueryReq.Status],
		Name:     resourceListQueryReq.Name,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	resourceListQueryResp := &dto.ResourceListQueryResp{
		Page:  resourceListQueryReq.Page,
		Size:  resourceListQueryReq.Size,
		Total: int(total),
	}
	for _, resource := range results {
		resourceListQueryResp.Resources = append(resourceListQueryResp.Resources, &dto.Resource{
			ResourceID:  resource.ResourceID,
			DomainID:    resource.DomainID,
			Name:        resource.Name,
			Description: resource.Description,
			Status:      convertor.ResourceStatus2DTO[resource.Status],
		})
	}

	dto.SuccessResp(c, resourceListQueryResp)
}
