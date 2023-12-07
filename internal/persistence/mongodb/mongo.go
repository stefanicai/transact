package mongodb

import (
	"context"
	"fmt"
	"github.com/stefanicai/transact/internal/model"
	"github.com/stefanicai/transact/internal/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

const (
	idKey = "id"
)

type transactionDao struct {
	client *mongo.Client
}

// Store persists a create.go
func (t *transactionDao) Store(ctx context.Context, tr *model.Transaction) error {
	tc := t.client.Database("test").Collection("transact")
	_, err := tc.InsertOne(ctx, *tr)
	if err != nil {
		slog.Error("failed inserting transaction", err)
		return err
	}
	return nil
}

func (t *transactionDao) Get(ctx context.Context, ID string) (*model.Transaction, error) {
	var result model.Transaction
	err := t.collection().FindOne(ctx, bson.D{{idKey, ID}}).Decode(&result)
	if err != nil {
		slog.Info("transaction not found", "id", ID, "error", err)
		return nil, err
	}

	return &result, nil
}

func (t *transactionDao) collection() *mongo.Collection {
	return t.client.Database("test").Collection("transact")
}

func MakeTransactionDao(client *mongo.Client) (persistence.TransactionDao, error) {
	return &transactionDao{
		client: client,
	}, nil
}

func CreateMongoClient(ctx context.Context, cfg Config) (*mongo.Client, error) {
	slog.Debug("creating Mongo DB client", "url", cfg.URL)
	var cred options.Credential

	//cred.AuthSource = YourAuthSource
	cred.Username = cfg.Username
	cred.Password = cfg.Password

	// set client options
	clientOptions := options.Client().ApplyURI(cfg.URL).SetAuth(cred)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
