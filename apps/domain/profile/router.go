package profile

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/middleware"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/controller"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/repositories/postgres"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/services/usecase"
)

type ProfileRouter interface {
	RegisterRoute()
}

type profileRouter struct {
	router *gin.RouterGroup
	middleware *middleware.MiddlewareGin
	controller *controller.ControllerAPI
}

func NewProfileRouter(route *gin.RouterGroup, db *sql.DB, middleware *middleware.MiddlewareGin) ProfileRouter {
	repo := postgres.NewProfileRepo(db)
	svc := usecase.NewProfileSvc(repo)
	controller := controller.NewControllerAPI(svc)

	return &profileRouter{
		router: route,
		middleware: middleware,
		controller: controller,
	}
}

func (p *profileRouter) RegisterRoute() {
	p.router.Use(p.middleware.ValidateAuth)
	p.router.GET("/profile", p.controller.UserProfile)
	p.router.POST("/search", p.controller.SearchProfile)
	p.router.GET("/list", p.controller.UserProfileList)
	p.router.GET("/follower", p.controller.UserFollowerList)
}