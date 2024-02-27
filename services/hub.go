package services

import (
	"log"

	"github.com/google/uuid"
)

type HubCOnfiguration func(ts *Hub) error
type Hub struct {
	OrderService    *OrderService
	ShippingService *ShippingService
}

func NewHub(cfgs ...HubCOnfiguration) (*Hub, error) {
	h := &Hub{}
	for _, cfg := range cfgs {
		err := cfg(h)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func WithOrderService(os *OrderService) HubCOnfiguration {
	return func(h *Hub) error {
		h.OrderService = os
		return nil
	}
}
func WithShippingService(ss *ShippingService) HubCOnfiguration {
	return func(h *Hub) error {
		h.ShippingService = ss
		return nil
	}
}

func (h *Hub) OrderAndShip(customerID uuid.UUID, productsIDs []uuid.UUID, shippingIDs []uuid.UUID) error {
	price, err := h.OrderService.CreateOrder(customerID, productsIDs)
	if err != nil {
		return err
	}
	statuses, err := h.ShippingService.CreateShipping(productsIDs, shippingIDs)
	if err != nil {
		return err
	}
	log.Printf("Customer made a bill of :%0.0f", price)
	for i, _ := range productsIDs {
		log.Printf("products %s shipped with shipping id %s has status  %s", productsIDs[i], shippingIDs[i], statuses[i])
	}
	return nil
}
