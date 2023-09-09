package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/invite/internal/adapter/http/handler"
	"github.com/hramov/invite/internal/config"
	"github.com/hramov/invite/pkg/database/postgres"
	"github.com/hramov/invite/pkg/database/types"
	"github.com/hramov/invite/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type Api struct {
	config *config.Config
	log    logger.Logger
	db     *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config, log logger.Logger) (*Api, error) {
	connectOptions := types.ConnectOptions{
		Host:            cfg.Db.Postgres.Host,
		Port:            cfg.Db.Postgres.Port,
		User:            cfg.Db.Postgres.User,
		Password:        cfg.Db.Postgres.Password,
		Database:        cfg.Db.Postgres.Database,
		SslMode:         "disable",
		MaxOpenCons:     50,
		MaxIdleCons:     10,
		ConnMaxIdleTime: 1 * time.Minute,
		ConnMaxLifetime: 5 * time.Minute,
	}

	db, err := postgres.New(ctx, connectOptions, log)
	if err != nil {
		return nil, err
	}
	return &Api{
		config: cfg,
		log:    log,
		db:     db,
	}, nil
}

func (a *Api) registerHandlers(r chi.Router) {
	h := handler.New(a.config, a.db, a.log)
	r.Route("/invite", h.Register)
}

func (a *Api) StartServer(ctx context.Context) {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})

	r.Use(c.Handler)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Route("/api", a.registerHandlers)

	go func() {
		a.log.Info("starting server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.App.Port), r); err != nil {
			a.log.Error(fmt.Sprintf("cannot start server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	a.log.Info(fmt.Sprintf("starting graceful shutdown for server"))
	err := a.StopServer()
	if err != nil {
		a.log.Error(fmt.Sprintf("cannot stop server"))
	}
}

func (a *Api) StopServer() error {
	a.db.Close()
	return nil
}
