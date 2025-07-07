package handler

import (
	"auth/biz/handler/service"
	"auth/biz/middleware/jwt"
	"auth/biz/model/dto"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// GetUserInfo
//
//	@Tags			user
//	@Summary		用户
//	@Description	用户
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer jwt"
//	@Success		200				{object}	dto.CommonResp{data=dto.LogoutResp}
//	@Router			/api/v1/user/info [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	payload := jwt.GetPayload(ctx)

	userInfo, bizErr := service.GetUserInfo(ctx, payload.UserID)
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	dto.SuccessResp(c, &dto.UserInfoQueryResp{
		Name:        userInfo.Name,
		Email:       userInfo.Email,
		Phone:       userInfo.Phone,
		Description: userInfo.Description,
	})
}
