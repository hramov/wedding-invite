package postgres

import (
	"context"
	"fmt"
	"github.com/hramov/invite/internal/adapter/http/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) SaveGuest(ctx context.Context, dto dto.GuestDto) (int, error) {
	sql := `
		INSERT INTO guest (name, email, presence, drink, food, music, transfer, accommodation, date_created)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
		RETURNING id;
	`

	params := []any{dto.Name, dto.Email, dto.Presence, dto.Drink, dto.Food, dto.Music, dto.Transfer, dto.Accommodation}

	row := r.conn.QueryRow(ctx, sql, params...)

	var guestId int
	err := row.Scan(&guestId)
	if err != nil {
		return 0, fmt.Errorf("cannot scan save guest result: %v", err)
	}

	return guestId, nil
}
