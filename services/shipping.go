package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
	"github.com/mallikarjunreddyD/DDD3/repositories/shipping"
	shipmem "github.com/mallikarjunreddyD/DDD3/repositories/shipping/memory"
)

type ShippingConfiguration func(ss *ShippingService) error

type ShippingService struct {
	shippings shipping.ShippingRepository
}

func NewShippingService(cfgs ...ShippingConfiguration) (*ShippingService, error) {
	ss := &ShippingService{}
	for _, cfg := range cfgs {
		err := cfg(ss)
		if err != nil {
			return nil, err
		}
	}
	return ss, nil
}

func withShippingRepository(sh shipping.ShippingRepository) ShippingConfiguration {
	return func(ss *ShippingService) error {
		ss.shippings = sh
		return nil
	}
}

func withMemoryShippingRepository() ShippingConfiguration {
	cr := shipmem.New()
	return withShippingRepository(cr)
}

func (ss *ShippingService) CreateShipping(productsIDs []uuid.UUID, shippingIDs []uuid.UUID) ([]string, error) {
	var shippings []aggregate.Shipping
	var status []string
	for _, id := range shippingIDs {
		p, err := ss.shippings.Get(id)
		if err != nil {
			return status, err
		}
		shippings = append(shippings, p)

	}
	for _, ship := range shippings {
		status = append(status, ship.GetStatus())
		log.Printf("shipping id is %s and status is %s", ship.GetID(), ship.GetStatus())
	}
	return status, nil
}
