package apps

import (
	"fmt"

	"github.com/ilhamnyto/twinyto/apps/router"
	"github.com/ilhamnyto/twinyto/pkg/database"
)

const (
	Router_GIN = "gin"
)

type RouterFactory struct {
	db *database.Database
}

func NewRouterFactory(db *database.Database) *RouterFactory {
	return &RouterFactory{
		db: db,
	}
}

type Router interface {
	BuildRoutes()
	Run()
}

func (r *RouterFactory) Create(routerType string, port string) (Router, error) {
	if routerType == Router_GIN {
		return router.NewRouterGin(r.db, port), nil
	}else {
		return nil, fmt.Errorf("router with type of %v is not available", routerType)
	}
}

type RouterExecutor struct {}

func NewRouterExecutor() *RouterExecutor {
	return &RouterExecutor{}
}

func (r *RouterExecutor) Execute(router Router) {
	router.BuildRoutes()
	router.Run()
}