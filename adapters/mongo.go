package adapters

import (
	"context"
	"fmt"
	"github.com/86soft/healthyro-recipes/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	dbName          = "healthyro-recipes"
	documentRecipes = "recipes"
)

type MongoStorage struct {
	client *mongo.Client
}

var _ domain.Repository = (*MongoStorage)(nil)

func NewMongoClient(uri string, timeoutInSec time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutInSec*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func NewMongoStorage(client *mongo.Client) *MongoStorage {
	return &MongoStorage{client: client}
}

func (m *MongoStorage) AddRecipe(ctx context.Context, r *domain.Recipe) error {
	dao := Recipe{
		Id:          r.ID(),
		CreatedAt:   time.Now().UTC(),
		Title:       r.Title(),
		Description: r.Description(),
	}
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)

	res, err := c.InsertOne(ctx, dao)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStorage) GetRecipe(ctx context.Context, id domain.RecipeID) (domain.Recipe, error) {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)
	dao := Recipe{}
	err := c.FindOne(ctx, bson.D{{"_id", id.GetID()}}).Decode(&dao)
	if err != nil {
		return domain.Recipe{}, err
	}
	return domain.UnmarshalRecipe(domain.NewRecipeID(dao.Id), dao.Title, dao.Description), nil
}

func (m *MongoStorage) GetRecipes(ctx context.Context) ([]domain.Recipe, error) {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)
	cursor, err := c.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var recipes []domain.Recipe
	for cursor.Next(ctx) {
		dao := Recipe{}
		err := cursor.Decode(&dao)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes,
			domain.UnmarshalRecipe(domain.NewRecipeID(dao.Id), dao.Title, dao.Description))
	}
	return recipes, nil
}

func (m *MongoStorage) UpdateRecipeTitle(ctx context.Context, id domain.RecipeID, title string) error {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)
	update := bson.D{{"$set", bson.D{{"title", title}}}}
	res, err := c.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount != 1 {
		//TODO: do smth
	}
	return nil
}

func (m *MongoStorage) UpdateRecipeDescription(ctx context.Context, id domain.RecipeID, description string) error {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)

	update := bson.D{{"$set", bson.D{{"description", description}}}}
	res, err := c.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		//TODO: do smth
	}
	return nil
}

func (m *MongoStorage) DeleteRecipe(ctx context.Context, id domain.RecipeID) error {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)

	_, err := c.DeleteOne(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStorage) AddRecipeResource(ctx context.Context, id domain.RecipeID, rsc *domain.Resource) error {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)

	r := Resource{
		Name:  rsc.Name,
		Kind:  rsc.Kind,
		Value: rsc.Value,
	}
	update := bson.D{{"$push", bson.D{{"resources", r}}}}
	_, err := c.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStorage) DeleteRecipeResource(ctx context.Context, id domain.RecipeID, name string) error {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)
	update := bson.M{"$pull": bson.M{"resources": bson.M{"name": name}}}
	_, err := c.UpdateByID(ctx, id, update)
	return err
}
