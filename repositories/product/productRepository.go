package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists: add failed")
	ErrProductDoesNotExists = errors.New("product does not exists: update failed")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	Get(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(id uuid.UUID) error
}
