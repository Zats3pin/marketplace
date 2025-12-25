package service

import (
	"MarketPlaceSolo/internal/domain"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *domain.User) error {
	query := `
		INSERT INTO user (name,Email)
		VALUES ($1,$2)
		RETURNING id, created_at
		`

	return r.db.QueryRow(
		query,
		u.Username,
		u.Email,
	).Scan(&u.ID, &u.CreatedAt)

}
