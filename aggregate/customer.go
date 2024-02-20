package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/entity"
	valueobjects "github.com/mallikarjunreddyD/DDD3/valueObjects"
)

var ErrInvalidCustomerName = errors.New(" Customer must have an valid name")

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobjects.Transaction
	addresses    []valueobjects.Address
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidCustomerName
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobjects.Transaction, 0),
		addresses:    make([]valueobjects.Address, 0),
	}, nil
}
