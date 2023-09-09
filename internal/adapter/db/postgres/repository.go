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
		INSERT INTO guest (name, email, presence, drink, food, music, transfer, accommodation, companion, date_created)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now())
		RETURNING id;
	`

	params := []any{dto.Name, dto.Email, dto.Presence, dto.Drink, dto.Food, dto.Music, dto.Transfer, dto.Accommodation, dto.Companion}

	row := r.conn.QueryRow(ctx, sql, params...)

	var guestId int
	err := row.Scan(&guestId)
	if err != nil {
		return 0, fmt.Errorf("cannot scan save guest result: %v", err)
	}

	return guestId, nil
}

func (r *Repository) GetCompanions(ctx context.Context) ([]dto.GuestDto, error) {
	sql := `
		SELECT max(id) as id, name, max(date_created) as date_created
		FROM guest
		WHERE presence = 'yes'
		GROUP BY name
	`

	var params []any

	rows, err := r.conn.Query(ctx, sql, params...)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch companions: %v", err)
	}

	var companions []dto.GuestDto

	for rows.Next() {
		var companion dto.GuestDto
		err = rows.Scan(&companion.Id, &companion.Name, &companion.DateCreated)
		if err != nil {
			return nil, fmt.Errorf("cannot scan companion: %v", err)
		}
		companions = append(companions, companion)
	}

	return companions, nil
}
