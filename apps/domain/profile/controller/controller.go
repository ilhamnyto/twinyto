package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/params"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/services"
)

type ControllerAPI struct {
	svc services.ProfileSvc
}

func NewControllerAPI(svc services.ProfileSvc) *ControllerAPI {
	return &ControllerAPI{
		svc: svc,
	}
}

func (c *ControllerAPI) UserProfile(ctx *gin.Context) {
	var (
		result *params.UserProfileResponse
		custErr *response.CustomError
	)

	userId := ctx.GetInt("UserID")
	username := ctx.Query("username")

	if username != "" {
		result, custErr = c.svc.UserProfile(ctx, username)
	}else {
		result, custErr = c.svc.MyProfile(ctx.Request.Context(), userId)
	}

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GeneralSuccessWithCustomMessageAndPayload("SUCCESS", result)
	
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) SearchProfile(ctx *gin.Context) {
	var req = new(params.UserProfileRequest)

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralErrorWithAdditionalInfo(err)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	result, custErr := c.svc.SearchProfile(ctx.Request.Context(), req)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessWithCustomMessageAndPayload("SUCCESS", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) UserProfileList(ctx *gin.Context) {
	result, custErr := c.svc.UserProfileList(ctx)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessWithCustomMessageAndPayload("SUCCESS", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) UserFollowerList(ctx *gin.Context) {
	userId := ctx.GetInt("UserID")

	result, custErr := c.svc.UserFollowerList(ctx.Request.Context(), userId)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessWithCustomMessageAndPayload("SUCCESS", result)
	ctx.JSON(resp.StatusCode, resp)
}