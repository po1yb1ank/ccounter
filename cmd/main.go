package main

import (
	"fmt"

	"github.com/po1yb1ank/ccounter/config"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/server"
	"github.com/po1yb1ank/ccounter/pkg/storage"
)

// @title Counter API
// @version 1.0
// @description Counter API service

// @host localhost:8888
// @BasePath /
// @query.collection.format multi
func main() {
	cfg, err := config.LoadConfig(config.DEFAULT_PATH)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	logger := logger.NewZapSugaredLogger()
	storage := storage.NewRedisStorage(cfg.Redis)
	server := server.NewGinServer(cfg.Server, logger, storage)

	if err := server.Run(); err != nil {
		logger.Error(err.Error())
	}
}
