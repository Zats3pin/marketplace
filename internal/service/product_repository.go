package service

import (
	"MarketPlaceSolo/internal/domain"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(p *domain.Product) error {
	query := `
		INSERT INTO products (name,price,stock)
		VALUES ($1,$2,$3)
		RETURNING id, created_at
		`

	return r.db.QueryRow(
		query,
		p.Name,
		p.Price,
		p.Stock,
	).Scan(&p.ID, &p.CreatedAt)

}
