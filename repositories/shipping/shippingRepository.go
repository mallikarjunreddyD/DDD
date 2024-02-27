package shipping

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD/aggregate"
)

var (
	ErrShippingNotFound      = errors.New("shipping not found")
	ErrShippingAlreadyExists = errors.New("shipping already exists: add failed")
	ErrShippingDoesNotExists = errors.New("shipping does not exists: update failed")
)

type ShippingRepository interface {
	GetAll() ([]aggregate.Shipping, error)
	Get(uuid.UUID) (aggregate.Shipping, error)
	Add(aggregate.Shipping) error
	Update(aggregate.Shipping) error
	Delete(id uuid.UUID) error
}
