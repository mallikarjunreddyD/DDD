package aggregate

import (
	"time"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/entity"
	valueobjects "github.com/mallikarjunreddyD/DDD3/valueObjects"
)

type Shipping struct {
	ship    *entity.Ship
	product *entity.Item
	address valueobjects.Address
}

func NeWShipping(mode string, deliveryDate time.Time, product *entity.Item, address valueobjects.Address) (Shipping, error) {
	return Shipping{
		ship: &entity.Ship{
			ID:           uuid.New(),
			Mode:         mode,
			Status:       "Ordered",
			DelivaryDate: deliveryDate,
		},
		product: product,
		address: address,
	}, nil
}

func (s Shipping) GetID() uuid.UUID {
	return s.ship.ID
}
func (s Shipping) SetID(id uuid.UUID) {
	if s.ship == nil {
		s.ship = &entity.Ship{}
	}
	s.ship.ID = id
}
func (s Shipping) GetStatus() string {
	return s.ship.Status
}
func (s Shipping) SetStatus(status string) {
	if s.ship == nil {
		s.ship = &entity.Ship{}
	}
	s.ship.Status = status
}
