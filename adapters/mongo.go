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
	documentTags    = "tags"
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
	createdAt := time.Now().UTC()

	daoRecipeResources := make([]Resource, 0, len(r.Resources))
	for _, res := range r.Resources {
		daoRecipeResources = append(daoRecipeResources, Resource{
			Document: Document{CreatedAt: createdAt},
			Id:       res.Id.GetID(),
			Name:     res.Name,
			Kind:     res.Kind,
			Value:    res.Value,
		})
	}

	daoRecipeTags := make([]RecipeTag, 0, len(r.Tags))
	for _, t := range r.Tags {
		daoRecipeTags = append(daoRecipeTags, RecipeTag{
			Id:   t.Id.GetID(),
			name: t.Name,
		})
	}

	dao := Recipe{
		Document: Document{
			CreatedAt: createdAt,
		},
		Id:          r.Id.GetID(),
		Title:       r.Title,
		Description: r.Description,
		Resources:   daoRecipeResources,
		Tags:        daoRecipeTags,
	}

	tags := make([]interface{}, 0, len(daoRecipeTags))
	for _, tag := range daoRecipeTags {
		tags = append(tags, Tag{
			Document: Document{CreatedAt: createdAt},
			ID:       tag.Id,
			Name:     tag.name,
			RecipeIDS: []string{
				r.Id.GetID(),
			},
		})
	}
	recipesCol := m.client.
		Database(dbName).
		Collection(documentRecipes)

	tagsCol := m.client.
		Database(dbName).
		Collection(documentTags)

	_, err := recipesCol.InsertOne(ctx, dao)
	if err != nil {
		return err
	}
	_, err = tagsCol.InsertMany(ctx, tags)
	return err
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
	return domain.Recipe{
		Id:          domain.NewRecipeID(dao.Id),
		Title:       "",
		Description: "",
		Resources:   nil,
		Tags:        nil,
	}, nil
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

		daoId := domain.NewRecipeID(dao.Id)
		resources := make([]domain.Resource, 0, len(dao.Resources))
		for _, r := range dao.Resources {
			resources = append(resources, domain.Resource{
				Id:    domain.ResourceID{Id: r.Id},
				Name:  r.Name,
				Kind:  r.Kind,
				Value: r.Value,
			})
		}

		tags := make([]domain.Tag, 0, len(dao.Tags))
		for _, d := range dao.Tags {
			tags = append(tags, domain.Tag{
				Id:       domain.NewTagID(d.Id),
				RecipeId: daoId,
				Name:     d.name,
			})
		}
		recipes = append(recipes,
			domain.Recipe{
				Id:          domain.NewRecipeID(dao.Id),
				Title:       dao.Title,
				Description: dao.Description,
				Resources:   resources,
				Tags:        tags,
			})
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
