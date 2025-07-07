package handler

import (
	"auth/biz/model/dto"
	"auth/biz/model/errs"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// CreatePermission
//
//	@Summary		创建权限
//	@Description	创建权限
//	@Tags			permission
//	@Accept			json
//	@Produce		json
//	@Param			permission	body		dto.PermissionCreateReq	true	"create permission request body"
//	@Success		200			{object}	dto.CommonResp{data=dto.PermissionCreateResp}
//	@Router			/api/v1/permission/create [POST]
func CreatePermission(ctx context.Context, c *app.RequestContext) {
	var permissionCreateReq dto.PermissionCreateReq
	if stdErr := c.BindAndValidate(&permissionCreateReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}
}
