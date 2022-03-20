package adapters

import (
	"context"
	"fmt"
	d "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *MongoStorage) CreateRecipe(ctx context.Context, recipe *d.Recipe) error {
	createdAt := time.Now().UTC()

	dbRecipeResources := make([]Resource, 0, len(recipe.Resources))
	mapToResources(createdAt, recipe.Resources, dbRecipeResources)

	dbRecipeTags := make([]RecipeTag, 0, len(recipe.Tags))
	mapToRecipeTags(recipe.Tags, dbRecipeTags)

	dbRecipe := Recipe{
		Document: Document{
			CreatedAt: createdAt,
		},
		ID:          recipe.ID.ID,
		Title:       recipe.Title,
		Description: recipe.Description,
		Resources:   dbRecipeResources,
		Tags:        dbRecipeTags,
	}

	dbTags := make([]any, 0, len(dbRecipeTags))
	mapToTags(createdAt, recipe, recipe.Tags, dbTags)

	recipesCol := m.ForCollection(CollectionRecipes)
	tagsCol := m.ForCollection(CollectionTags)

	_, errOrNil := recipesCol.InsertOne(ctx, dbRecipe)
	if errOrNil != nil {
		return errOrNil
	}
	_, errOrNil = tagsCol.InsertMany(ctx, dbTags)
	return errOrNil
}

func (m *MongoStorage) AddRecipeResource(ctx context.Context, id d.ID[d.Recipe], r *d.Resource) error {
	createdAt := time.Now().UTC()
	recipeColl := m.ForCollection(CollectionRecipes)

	update := bson.M{"$push": bson.M{"resources": Resource{
		Document: Document{CreatedAt: createdAt},
		ID:       r.ID.ID,
		Name:     r.Name,
		Kind:     r.Kind,
		Value:    r.Value,
	}}}
	_, errOrNil := recipeColl.UpdateByID(ctx, id.ID, update)
	if errOrNil != nil {
		return errOrNil
	}
	return errOrNil
}

func (m *MongoStorage) GetRecipe(ctx context.Context, id d.ID[d.Recipe]) (d.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)

	dbRecipe := Recipe{}
	err := c.FindOne(ctx, bson.D{{"_id", id.ID}}).Decode(&dbRecipe)
	if err != nil {
		return d.Recipe{}, err
	}
	rResources := make([]d.Resource, 0, len(dbRecipe.Resources))
	mapFromResources(dbRecipe.Resources, rResources)

	rTags := make([]d.Tag, 0, len(dbRecipe.Tags))
	mapFromRecipeTags(id, dbRecipe.Tags, rTags)

	return d.Recipe{
		ID:          id,
		Title:       dbRecipe.Title,
		Description: dbRecipe.Description,
		Resources:   rResources,
		Tags:        rTags,
	}, nil
}

func (m *MongoStorage) ListRecipes(ctx context.Context) ([]d.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)

	cursor, err := c.Find(ctx, bson.D{{}}) // should consider count, but we need pagination in future anyway
	if err != nil {
		return nil, err
	}
	var recipes []d.Recipe

	for cursor.Next(ctx) {
		dbRecipe := Recipe{}
		err := cursor.Decode(&dbRecipe)
		if err != nil {
			return nil, err
		}
		id := d.FromStringID[d.Recipe](dbRecipe.ID)

		rResources := make([]d.Resource, 0, len(dbRecipe.Resources))
		mapFromResources(dbRecipe.Resources, rResources)

		tags := make([]d.Tag, 0, len(dbRecipe.Tags))
		mapFromRecipeTags(id, dbRecipe.Tags, tags)

		recipes = append(recipes,
			d.Recipe{
				ID:          id,
				Title:       dbRecipe.Title,
				Description: dbRecipe.Description,
				Resources:   rResources,
				Tags:        tags,
			})
	}
	return recipes, nil
}

func (m *MongoStorage) UpdateRecipeTitle(ctx context.Context, id d.ID[d.Recipe], title string) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.D{{"$set", bson.D{{"title", title}}}}
	res, err := c.UpdateByID(ctx, id.ID, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return &d.OnDBUpdateError{
			ID:      id.ID,
			Details: fmt.Sprintf("ModifiedCount is different than 1, count: %v", res.ModifiedCount),
		}
	}

	return nil
}

func (m *MongoStorage) UpdateRecipeDescription(ctx context.Context, id d.ID[d.Recipe], description string) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.D{{"$set", bson.D{{"description", description}}}}
	res, err := c.UpdateByID(ctx, id.ID, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return &d.OnDBUpdateError{
			ID:      id.ID,
			Details: fmt.Sprintf("ModifiedCount is different than 1, count: %v", res.ModifiedCount),
		}
	}

	return nil
}

func (m *MongoStorage) DeleteRecipe(ctx context.Context, id d.ID[d.Recipe]) error {
	c := m.ForCollection(CollectionRecipes)
	_, errOrNil := c.DeleteOne(ctx, id.ID)
	return errOrNil
}

func (m *MongoStorage) DeleteRecipeResource(ctx context.Context, recipeID d.ID[d.Recipe], resourceID d.ID[d.Resource]) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.M{"$pull": bson.M{"resources": bson.M{"_id": resourceID.ID}}}
	_, errOrNil := c.UpdateByID(ctx, recipeID.ID, update)
	return errOrNil
}
