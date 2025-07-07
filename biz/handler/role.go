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

// UpdateRole
//
//	@Summary		更新角色
//	@Description	更新角色
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Param			role_id	query		string				true	"update role"
//	@Param			role	body		dto.RoleUpdateReq	true	"update role"
//	@Success		200		{object}	dto.CommonResp{data=dto.RoleUpdateResp}
//	@Router			/api/v1/role/update [POST]
func UpdateRole(ctx context.Context, c *app.RequestContext) {
	var roleUpdateReq dto.RoleUpdateReq
	if stdErr := c.BindAndValidate(&roleUpdateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	if bizErr := service.UpdateRole(ctx, &service.RoleUpdateParam{
		RoleID:      roleUpdateReq.RoleID,
		Status:      convertor.RoleStatus2DO[roleUpdateReq.Status],
		Name:        roleUpdateReq.Name,
		Description: roleUpdateReq.Description,
	}); bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	dto.SuccessResp(c, &dto.RoleUpdateResp{})
}

// DeleteRole
//
//	@Summary		删除角色
//	@Description	删除角色
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Param			role_id	query		string	true	"delete role"
//	@Success		200		{object}	dto.CommonResp{data=dto.RoleDeleteResp}
//	@Router			/api/v1/role/delete [POST]
func DeleteRole(ctx context.Context, c *app.RequestContext) {
	var roleDeleteReq dto.RoleDeleteReq
	if stdErr := c.BindAndValidate(&roleDeleteReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	if bizErr := service.DeleteRole(ctx, roleDeleteReq.RoleID); bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.RoleDeleteResp{})
}

// QueryRoleDetail
//
//	@Summary		查询角色详情
//	@Description	查询角色详情
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.RoleDetailQueryReq	true	"query role"
//	@Success		200	{object}	dto.CommonResp{data=dto.RoleDetailQueryResp}
//	@Router			/api/v1/role/query [GET]
func QueryRoleDetail(ctx context.Context, c *app.RequestContext) {
	var roleDetailQueryReq dto.RoleDetailQueryReq
	if stdErr := c.BindAndValidate(&roleDetailQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	result, bizErr := service.QueryRoleDetail(ctx, roleDetailQueryReq.RoleID)
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.RoleDetailQueryResp{
		Role: dto.Role{
			RoleID:      result.RoleID,
			DomainID:    result.DomainID,
			ParentID:    result.ParentID,
			Name:        result.Name,
			Status:      convertor.RoleStatus2DTO[result.Status],
			Description: result.Description,
		},
	})
}
