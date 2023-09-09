package adapter

import (
	"context"
	"fmt"
	v1 "github.com/hramov/invite/internal/adapter/http"
	"github.com/hramov/invite/internal/config"
	"github.com/hramov/invite/pkg/logger"
)

func StartApp(ctx context.Context, cfg *config.Config) {
	log := logger.New(cfg.App.Name, logger.Debug)
	server, err := v1.New(ctx, cfg, log)
	if err != nil {
		log.Error(fmt.Sprintf("cannot start server: %v, exiting...", err))
		return
	}
	server.StartServer(ctx)
}
