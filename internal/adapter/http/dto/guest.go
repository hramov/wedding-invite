package dto

import "time"

type GuestDto struct {
	Id            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Email         string    `json:"email,omitempty"`
	Presence      string    `json:"presence,omitempty"`
	Drink         []string  `json:"drink,omitempty"`
	Food          []string  `json:"food,omitempty"`
	Music         string    `json:"music,omitempty"`
	Transfer      string    `json:"transfer,omitempty"`
	Accommodation string    `json:"accommodation,omitempty"`
	Companion     int       `json:"companion,omitempty"`
	DateCreated   time.Time `json:"date_created"`
}
