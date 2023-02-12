package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/pkg/token"
)

type MiddlewareGin struct {}

func NewMiddlewareGin() *MiddlewareGin {
	return &MiddlewareGin{}
}

func (m *MiddlewareGin) ValidateAuth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	t := strings.Split(tokenString, "Bearer ")

	if len(t) < 2 {
		resp := response.GeneralErrorWithAdditionalInfo("invalid len of token.")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	payload, err := token.ValidateToken(t[1])

	if err != nil {
		resp := response.GeneralErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	now := time.Now()

	if now.After(payload.Expired) {
		resp := response.GeneralErrorWithAdditionalInfo("token is expired")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	ctx.Set("UserID", payload.UserId)

	ctx.Next()
}