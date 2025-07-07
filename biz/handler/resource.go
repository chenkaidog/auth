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

// UpdateResource
//
//	@Summary		更新资源
//	@Description	更新资源
//	@Tags			resource
//	@Accept			json
//	@Produce		json
//	@Param			resource_id	query		string					true	"delete resource request body"
//	@Param			resource	body		dto.ResourceUpdateReq	true	"update resource request body"
//	@Success		200			{object}	dto.CommonResp{data=dto.ResourceUpdateResp}
//	@Router			/api/v1/resource/update [POST]
func UpdateResource(ctx context.Context, c *app.RequestContext) {
	var resourceUpdateReq dto.ResourceUpdateReq
	if stdErr := c.BindAndValidate(&resourceUpdateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	if bizErr := service.UpdateResource(ctx, &service.ResourceUpdateParam{
		ResourceID:  resourceUpdateReq.ResourceID,
		Name:        resourceUpdateReq.Name,
		Description: resourceUpdateReq.Description,
		Status:      convertor.ResourceStatus2DO[resourceUpdateReq.Status],
	}); bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.ResourceUpdateResp{})
}

// DeleteResource
//
//	@Summary		删除资源
//	@Description	删除资源
//	@Tags			resource
//	@Accept			json
//	@Produce		json
//	@Param			resource_id	query		string	true	"delete resource"
//	@Success		200			{object}	dto.CommonResp{data=dto.ResourceDeleteResp}
//	@Router			/api/v1/resource/delete [POST]
func DeleteResource(ctx context.Context, c *app.RequestContext) {
	var resourceDeleteReq dto.ResourceDeleteReq
	if stdErr := c.BindAndValidate(&resourceDeleteReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	if bizErr := service.DeleteResource(ctx, resourceDeleteReq.ResourceID); bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.ResourceDeleteResp{})
}

// QueryResourceDetail
//
//	@Summary		查询资源详情
//	@Description	查询资源详情
//	@Tags			resource
//	@Accept			json
//	@Produce		json
//	@Param			req	query		dto.ResourceDetailQueryReq	true	"query resource"
//	@Success		200	{object}	dto.CommonResp{data=dto.ResourceDetailQueryResp}
//	@Router			/api/v1/resource/query [GET]
func QueryResourceDetail(ctx context.Context, c *app.RequestContext) {
	var resourceDetailQueryReq dto.ResourceDetailQueryReq
	if stdErr := c.BindAndValidate(&resourceDetailQueryReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}
	resourceDetail, bizErr := service.QueryResourceDetail(ctx, resourceDetailQueryReq.ResourceID)
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}
	dto.SuccessResp(c, &dto.ResourceDetailQueryResp{
		Resource: dto.Resource{
			ResourceID:  resourceDetail.ResourceID,
			DomainID:    resourceDetail.DomainID,
			Name:        resourceDetail.Name,
			Description: resourceDetail.Description,
			Status:      convertor.ResourceStatus2DTO[resourceDetail.Status],
		},
	})
}
