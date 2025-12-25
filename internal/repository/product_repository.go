package repository

import "MarketPlaceSolo/internal/domain"

type ProductRepository interface {
	Create(product *domain.Product) error
	GetById(id int64) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(id int64, product *domain.Product) error
	Delete(id int64) error
}
