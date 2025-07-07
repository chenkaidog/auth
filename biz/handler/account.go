package handler

import (
	"auth/biz/handler/service"
	"auth/biz/middleware/jwt"
	"auth/biz/middleware/session"
	"auth/biz/model/dto"
	"auth/biz/model/errs"
	"auth/biz/util/random"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

// Login 用户登录接口
//
//	@Tags			account
//	@Summary		用户登录接口
//	@Description	用户登录接口
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.LoginReq	true	"login request body"
//	@Success		200	{object}	dto.CommonResp{data=dto.LoginResp}
//	@Header			200	{string}	set-cookie	"cookie"
//	@Router			/api/v1/account/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var loginReq dto.LoginReq
	if stdErr := c.BindAndValidate(&loginReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	loginRes, bizErr := service.Login(ctx, service.LoginParam{
		Username: loginReq.Username,
		Password: loginReq.Password,
	})
	if bizErr != nil {
		dto.FailResp(c, bizErr)
		return
	}

	sess := sessions.Default(c)
	sess.Set(random.RandStr(32), 1) // set a random key to overwrite the old session
	sess.Clear()
	if err := sess.Save(); err != nil {
		hlog.CtxErrorf(ctx, "save session err: %v", err)
		dto.FailResp(c, errs.ServerError)
		return
	}

	jwtStr, expAt, err := jwt.GenerateToken(ctx, c, jwt.Payload{
		UserID: loginRes.UserID,
		Name:   loginRes.Name,
	}, sess.ID())
	if err != nil {
		dto.FailResp(c, errs.ServerError)
		return
	}

	dto.SuccessResp(c, &dto.LoginResp{
		ExpiresAt:   expAt,
		AccessToken: jwtStr,
	})
}

// Logout 用户登出接口
//
//	@Tags			account
//	@Summary		用户登出接口
//	@Description	用户登出接口
//	@Accept			json
//	@Produce		json
//	@Param			req				body		dto.LogoutReq	true	"logout request body"
//	@Param			Authorization	header		string			true	"Bearer jwt"
//	@Success		200				{object}	dto.CommonResp{data=dto.LogoutResp}
//	@Header			200				{string}	set-cookie	"cookie"
//	@Router			/api/v1/account/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var logoutReq dto.LogoutReq
	if stdErr := c.BindAndValidate(&logoutReq); stdErr != nil {
		hlog.CtxInfof(ctx, "BindAndValidate fail, %v", stdErr)
		dto.AbortWithErr(c, errs.ParamError, http.StatusBadRequest)
		return
	}

	_ = session.Remove(c)
	_ = jwt.RemoveToken(ctx)

	dto.SuccessResp(c, &dto.LogoutResp{})
}
