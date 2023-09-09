package postgres

import (
	"github.com/hramov/invite/internal/adapter/http/dto"
	"time"
)

type Guest struct {
	Id            int
	Name          string
	Email         string
	Presence      string
	Drink         []string
	Food          []string
	Music         string
	Transfer      string
	Accommodation string
	DateCreated   time.Time
}

func (u *Guest) Map() dto.GuestDto {
	return dto.GuestDto{
		Id:            u.Id,
		Name:          u.Name,
		Email:         u.Email,
		Presence:      u.Presence,
		Drink:         u.Drink,
		Food:          u.Food,
		Music:         u.Music,
		Transfer:      u.Transfer,
		Accommodation: u.Accommodation,
		DateCreated:   u.DateCreated,
	}
}
