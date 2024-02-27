package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mallikarjunreddyD/DDD3/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRespository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, connectionString string) (*MongoRespository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("DDD3")
	customers := db.Collection("customers")

	return &MongoRespository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongoRespository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}
	return c.ToAggregate(), nil
}

func (mr *MongoRespository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	in := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
func (mr *MongoRespository) Update(c aggregate.Customer) error {
	panic("to be added later")
}
