package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/middleware"
	"github.com/ilhamnyto/twinyto/apps/domain/post/repositories/postgres"
	"github.com/ilhamnyto/twinyto/apps/domain/post/services"
	"github.com/ilhamnyto/twinyto/apps/domain/post/services/usecase"
)

type ControllerAPI interface {
	RegisterRoutes()
}

type controllerAPI struct {
	route *gin.RouterGroup
	middleware *middleware.MiddlewareGin
	svc services.PostSvc
}

func NewControllerAPI(route *gin.RouterGroup, db *sql.DB, middleware *middleware.MiddlewareGin) ControllerAPI {
	repo := postgres.NewPostRepo(db)
	services := usecase.NewPostSvc(repo)
	return &controllerAPI{
		route: route,
		svc: services,
		middleware: middleware,
	}
}

func (c *controllerAPI) RegisterRoutes() {
}