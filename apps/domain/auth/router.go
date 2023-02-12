package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/middleware"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/controller"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/repositories/postgres"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/services/usecase"
)

type authRouter struct {
	route *gin.RouterGroup
	middleware *middleware.MiddlewareGin
	controller *controller.ControllerAPI
}

type AuthRouter interface {
	RegisterRoute()
}

func NewAuthRouter(route *gin.RouterGroup, db *sql.DB, middleware *middleware.MiddlewareGin) AuthRouter {
	repo := postgres.NewAuthRepo(db)
	svc := usecase.NewAuthSvc(repo)
	controller := controller.NewControllerAPI(svc)
	return &authRouter{
		route: route,
		middleware: middleware,
		controller: controller,
	}
}

func (a *authRouter) RegisterRoute() {
	a.route.POST("/login", a.controller.Login)
	a.route.POST("/register", a.controller.Register)
	a.route.POST("/reset-password", a.controller.ResetPassword)
}