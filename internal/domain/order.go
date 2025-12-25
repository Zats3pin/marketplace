package domain

import "time"

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Total     int64     `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created"`
}
