package server

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/po1yb1ank/ccounter/config"
	docs "github.com/po1yb1ank/ccounter/docs"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key/decrement"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key/increment"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key/reset"
	"github.com/po1yb1ank/ccounter/pkg/handlers/subscribe"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

type IServer interface {
	Run() error
}

type GinServer struct {
	addr string

	router  *gin.Engine
	storage storage.IStorage
	logger  logger.ILogger
	watcher watcher.IWatcher
}

func NewGinServer(
	config config.Server,
	logger logger.ILogger,
	storage storage.IStorage,
) IServer {
	router := gin.Default()
	addr := net.JoinHostPort(config.Host, config.Port)
	watcher := watcher.NewWSWatcher()

	server := &GinServer{
		router:  router,
		addr:    addr,
		logger:  logger,
		watcher: watcher,
		storage: storage,
	}

	server.registerHandlers()

	return server
}

func (g *GinServer) registerHandlers() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = g.addr

	g.router.GET("/:key", key.Get(g.storage, g.logger))

	g.router.POST("/:key/reset", reset.Post(g.storage, g.logger, g.watcher))
	g.router.POST("/:key/increment", increment.Post(g.storage, g.logger, g.watcher))
	g.router.POST("/:key/decrement", decrement.Post(g.storage, g.logger, g.watcher))

	g.router.GET("/subscribe", subscribe.Ws(g.storage, g.logger, g.watcher))

	g.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (g *GinServer) Run() error {
	g.logger.Info(g.getInfo())

	if err := g.router.Run(g.addr); err != nil {
		return err
	}

	return nil
}

func (g *GinServer) getInfo() string {
	return fmt.Sprintf("Starting server on: %v", g.addr)
}
