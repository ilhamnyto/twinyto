package follow

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/middleware"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/controller"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/repositories/postgres"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/services/usecase"
)

type FollowRouter interface {
	RegisterRoute()
}

type followRouter struct {
	router *gin.RouterGroup
	middleware *middleware.MiddlewareGin
	controller *controller.ControllerAPI
}

func NewFollowRouter(router *gin.RouterGroup, db *sql.DB, middleware *middleware.MiddlewareGin) FollowRouter {
	repo := postgres.NewFollowRepo(db)
	svc := usecase.NewFollowSvc(repo)
	controller := controller.NewControllerAPI(svc)

	return &followRouter{
		router: router,
		middleware: middleware,
		controller: controller,
	}
}

func (f *followRouter) RegisterRoute() {
	f.router.Use(f.middleware.ValidateAuth)
	f.router.POST("/follow", f.controller.Follow)
	f.router.POST("/unfollow", f.controller.Unfollow)
}