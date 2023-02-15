package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamnyto/twinyto/apps/commons/middleware"
	"github.com/ilhamnyto/twinyto/apps/domain/auth"
	"github.com/ilhamnyto/twinyto/apps/domain/profile"
	"github.com/ilhamnyto/twinyto/pkg/database"
)

type Gin struct {
	db *database.Database
	router *gin.Engine
	port string
	middleware *middleware.MiddlewareGin
}

func NewRouterGin(db *database.Database, port string) *Gin {
	router := gin.Default()
	middleware := middleware.NewMiddlewareGin()
	return &Gin{
		db: db,
		router: router,
		port: port,
		middleware: middleware,
	}
}

func (g *Gin) BuildRoutes() {
	g.router.Use(CORS)

	v1 := g.router.Group("api/v1")

	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {"message": "Hello twinyto."})
	})

	authPath := v1.Group("/auth")
	authRouter := auth.NewAuthRouter(authPath, g.db.DbSQL, g.middleware)
	authRouter.RegisterRoute()

	profilePath := v1.Group("/user")
	profileRouter := profile.NewProfileRouter(profilePath, g.db.DbSQL, g.middleware)
	profileRouter.RegisterRoute()
}

func (g *Gin) Run() {
	g.router.Run(g.port)
}

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Headers","Content-Type, Content-Length, X-Requested-With, X-CSRF-Token, Cache-Control, Accept-Encoding, Authorization, accept, origin")
	ctx.Header("Access-Control-Allow-Origin","*")
	ctx.Header("Access-Control-Allow-Credentials","true")
	ctx.Header("Access-Control-Allow-Methods","GET, POST, DELETE, PUT, OPTIONS")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}