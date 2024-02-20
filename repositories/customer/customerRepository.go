package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
)

var (
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrCustomerAlreadyExists = errors.New("customer already exists: failed to add customer")
	ErrCustomerDoesNotExists = errors.New("customer does not exists: update failed")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
