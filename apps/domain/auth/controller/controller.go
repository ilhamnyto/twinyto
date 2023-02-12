package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/params"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/services"
)

type ControllerAPI struct {
	svc services.AuthSvc
}

func NewControllerAPI(svc services.AuthSvc) *ControllerAPI {
	return &ControllerAPI{
		svc: svc,
	}
}

func (c *ControllerAPI) Login(ctx *gin.Context) {
	var req = new(params.UserLoginRequest)

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	result, custERr := c.svc.Login(ctx.Request.Context(), req)

	if custERr != nil {
		ctx.AbortWithStatusJSON(custERr.StatusCode, custERr)
		return
	}

	resp := response.GeneralSuccessWithCustomMessageAndPayload("LOGIN SUCCESS", result)

	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) Register(ctx *gin.Context) {
	var req = new(params.UserRegisterRequest)

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	custErr := c.svc.Register(ctx.Request.Context(), req)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) ResetPassword(ctx *gin.Context) {
	var req = new(params.UserPasswordResetRequest)

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	userId := ctx.GetInt("UserID")

	custErr := c.svc.ResetPassword(ctx.Request.Context(), req, userId)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}