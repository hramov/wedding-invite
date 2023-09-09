package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/invite/internal/adapter/db/postgres"
	"github.com/hramov/invite/internal/adapter/http/dto"
	"github.com/hramov/invite/internal/config"
	"github.com/hramov/invite/pkg/logger"
	"github.com/hramov/invite/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type Handler struct {
	cfg        *config.Config
	repository *postgres.Repository
	log        logger.Logger
}

func New(cfg *config.Config, db *pgxpool.Pool, log logger.Logger) *Handler {
	repo := postgres.NewRepository(db)
	return &Handler{cfg, repo, log}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/add-guest", h.addGuest)
	r.Get("/check", h.check)
}

func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(http.StatusOK, "Healthy\n", w)
}

func (h *Handler) addGuest(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[dto.GuestDto](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Errorf("cannot parse body: %v", err).Error(), w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	guestId, err := h.repository.SaveGuest(ctx, body)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Errorf("cannot save guest: %v", err).Error(), w)
		return
	}

	utils.SendResponse(http.StatusCreated, guestId, w)
}
