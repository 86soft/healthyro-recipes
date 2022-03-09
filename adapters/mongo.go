package adapters

import (
	"context"
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

func NewMongoClient(uri string, timeoutInSec time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutInSec*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}

func NewMongoStorage(client *mongo.Client) *MongoStorage {
	return &MongoStorage{client: client}
}

func (m *MongoStorage) AddRecipe(ctx context.Context, r *domain.Recipe) error {
	dao := recipe{
		id:          r.ID(),
		createdAt:   time.Now().UTC(),
		title:       r.Title(),
		description: r.Description(),
	}
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)

	_, err := c.InsertOne(ctx, dao)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStorage) GetRecipe(ctx context.Context, id domain.RecipeID) (domain.Recipe, error) {
	c := m.client.
		Database(dbName).
		Collection(documentRecipes)
	dao := recipe{}
	err := c.FindOne(ctx, bson.D{{"_id", id.GetID()}}).Decode(&dao)
	if err != nil {
		return domain.Recipe{}, err
	}
	return domain.UnmarshalRecipe(domain.NewRecipeID(dao.id), dao.title, dao.description), nil
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
		dao := recipe{}
		err := cursor.Decode(&dao)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes,
			domain.UnmarshalRecipe(domain.NewRecipeID(dao.id), dao.title, dao.description))
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

	r := resource{
		name:  rsc.Name(),
		kind:  rsc.Kind(),
		value: rsc.Value(),
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
