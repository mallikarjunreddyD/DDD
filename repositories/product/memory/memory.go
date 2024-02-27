package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD/aggregate"
	"github.com/mallikarjunreddyD/DDD/repositories/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}
func (mpr *MemoryProductRepository) Get(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}
func (mpr *MemoryProductRepository) Add(p aggregate.Product) error {
	if mpr.products == nil {
		mpr.products = make(map[uuid.UUID]aggregate.Product)
	}
	if _, ok := mpr.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}
	mpr.Lock()
	mpr.products[p.GetID()] = p
	mpr.Unlock()
	return nil
}
func (mpr *MemoryProductRepository) Update(p aggregate.Product) error {
	if _, ok := mpr.products[p.GetID()]; !ok {
		return product.ErrProductDoesNotExists
	}
	mpr.Lock()
	mpr.products[p.GetID()] = p
	mpr.Unlock()
	return nil
}
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductDoesNotExists
	}
	delete(mpr.products, id)
	return nil
}
