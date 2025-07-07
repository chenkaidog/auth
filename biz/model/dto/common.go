package dto

import (
	"auth/biz/model/errs"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type CommonResp struct {
	Success bool        `json:"success"`
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResp(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, &CommonResp{
		Success: true,
		Code:    errs.Success.Code(),
		Message: errs.Success.Msg(),
		Data:    data,
	})
}

func FailResp(c *app.RequestContext, bizErr errs.Error) {
	c.JSON(http.StatusOK, &CommonResp{
		Success: false,
		Code:    bizErr.Code(),
		Message: bizErr.Msg(),
	})
}

func AbortWithErr(c *app.RequestContext, bizErr errs.Error, httpCode int) {
	c.AbortWithStatusJSON(httpCode, &CommonResp{
		Success: false,
		Code:    bizErr.Code(),
		Message: bizErr.Msg(),
	})
}
