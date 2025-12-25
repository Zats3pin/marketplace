package repository

import "MarketPlaceSolo/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	List() ([]*domain.User, error)
	Update(id int64, user *domain.User) error
	Delete(id int64) error
}
