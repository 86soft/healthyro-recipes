package adapters

import (
	"context"
	"github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	dbName            = "healthyro-recipes"
	CollectionRecipes = "recipes"
	CollectionTags    = "tags"
)

type MongoStorage struct {
	client *mongo.Client
}

var _ core.Store = (*MongoStorage)(nil)

func NewMongoClient(uri string, timeoutInSec time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutInSec*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoStorage(client *mongo.Client) *MongoStorage {
	return &MongoStorage{client: client}
}

func (m *MongoStorage) ForCollection(name string) *mongo.Collection {
	return m.client.
		Database(dbName).
		Collection(name)
}
