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
