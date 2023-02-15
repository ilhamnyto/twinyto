package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/params"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/services"
)

type ControllerAPI struct {
	svc services.FollowSvc
}

func NewControllerAPI(svc services.FollowSvc) *ControllerAPI {
	return &ControllerAPI{
		svc: svc,
	}
}

func (c *ControllerAPI) Follow(ctx *gin.Context) {
	var req = new(params.UserFollowRequest)
	userId := ctx.GetInt("UserID")
	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralErrorWithAdditionalInfo(err)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	req.UserId = userId
	custErr := c.svc.Follow(ctx.Request.Context(), req)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) Unfollow(ctx *gin.Context) {
	var req = new(params.UserUnfollowRequest)
	userId := ctx.GetInt("UserID")
	err := ctx.ShouldBindJSON(req)

	if err != nil {
		resp := response.GeneralErrorWithAdditionalInfo(err)
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	req.UserId = userId
	custErr := c.svc.Unfollow(ctx.Request.Context(), req)

	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}