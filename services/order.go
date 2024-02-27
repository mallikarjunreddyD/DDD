package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD/aggregate"
	"github.com/mallikarjunreddyD/DDD/repositories/customer"
	"github.com/mallikarjunreddyD/DDD/repositories/customer/memory"
	"github.com/mallikarjunreddyD/DDD/repositories/customer/mongo"
	"github.com/mallikarjunreddyD/DDD/repositories/product"
	prodmem "github.com/mallikarjunreddyD/DDD/repositories/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func withCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func withMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return withCustomerRepository(cr)
}
func withMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}

}
func withProductRepository(cr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.products = cr
		return nil
	}
}

func withMemoryProductRepository() OrderConfiguration {
	cr := prodmem.New()
	return withProductRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, nil
	}
	var products []aggregate.Product
	var totalPrice float64

	for _, id := range productsIDs {
		p, err := o.products.Get(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		totalPrice += p.GetPrice()
	}
	log.Print(products)
	log.Printf("Customer: %s has ordered %d products for a toal of %f rupees", c.GetID(), len(products), totalPrice)
	return totalPrice, nil
}
