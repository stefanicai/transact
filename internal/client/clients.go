package client

import (
	"context"
	"github.com/stefanicai/transact/internal/config"
	"github.com/stefanicai/transact/internal/persistence/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Clients handles the creation and destruction of client objects at the start and end of the application
// Note: we could also just create the DAO at this point and pass it on. The reason why we don't do that is because it is possible that we'd want to use the client in other functionality.
// But specifically for this exercise, we could just create the DAO here and hide the entire database dependent code within that module.
type Clients interface {
	GetMongo() *mongo.Client
	Disconnect(ctx context.Context) error
}

type clients struct {
	mongoClient *mongo.Client
}

func (c *clients) GetMongo() *mongo.Client {
	return c.mongoClient
}

func (c *clients) Disconnect(ctx context.Context) error {
	return c.mongoClient.Disconnect(ctx)
}

// MakeClients create clients used by services
func MakeClients(ctx context.Context, cfg config.Config) (Clients, error) {
	var c clients
	if !cfg.Mongo.UseMock {
		mc, err := mongodb.CreateMongoClient(ctx, cfg.Mongo)
		if err != nil {
			return nil, err
		}
		c.mongoClient = mc
	}
	return &c, nil
}
