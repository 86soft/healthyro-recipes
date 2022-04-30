package adapters

import (
	"context"
	"fmt"
	c "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *MongoStorage) CreateTag(ctx context.Context, name string) (c.ID[c.Tag], error) {
	createdAt := time.Now().UTC()

	tagColl := m.ForCollection(CollectionTags)

	id := c.CreateID[c.Tag]()
	_, err := tagColl.InsertOne(ctx, Tag{
		Document: Document{CreatedAt: createdAt},
		Name:     name,
	})
	return id, err
}

func (m *MongoStorage) AddTagToRecipe(ctx context.Context, recipeID c.ID[c.Recipe], t *c.Tag) error {
	recipeColl := m.ForCollection(CollectionRecipes)

	update := bson.M{"$push": bson.M{"tags": t.Name}}

	_, err := recipeColl.UpdateByID(ctx, recipeID.Value, update)
	if err != nil {
		return fmt.Errorf("recipe: UpdateByID: %w", err)
	}

	tagColl := m.ForCollection(CollectionTags)
	update = bson.M{"$push": bson.M{"recipeIds": recipeID.Value}}
	_, err = tagColl.UpdateByID(ctx, t.Name, update)
	if err != nil {
		return fmt.Errorf("tagColl.UpdateByID: %w", err)
	}
	return nil
}

func (m *MongoStorage) RemoveTagFromRecipe(ctx context.Context, recipeID c.ID[c.Recipe], tagID c.ID[c.Tag]) error {
	recipeColl := m.ForCollection(CollectionRecipes)

	removeTagFromRecipe := bson.M{"$pull": bson.M{"tags": bson.M{"_id": tagID.Value}}}
	_, err := recipeColl.UpdateByID(ctx, recipeID.Value, removeTagFromRecipe)
	if err != nil {
		return fmt.Errorf("recipeColl.UpdateByID: %w", err)
	}
	return nil
}

func (m *MongoStorage) CreateTags(ctx context.Context, tags []c.Tag) error {
	createdAt := time.Now().UTC()

	tagColl := m.ForCollection(CollectionTags)

	dbTags := make([]any, 0, len(tags))
	for _, tag := range tags {

		dbTags = append(dbTags, Tag{
			Document:   Document{CreatedAt: createdAt},
			Name:       tag.Name,
			RecipesIDs: []string{},
		})
	}
	_, err := tagColl.InsertMany(ctx, dbTags)
	if err != nil {

	}
	return err
}

func (m *MongoStorage) AddRecipeToTags(ctx context.Context, recipeID c.ID[c.Recipe], tags []c.Tag) error {
	now := time.Now().UTC()
	tagColl := m.ForCollection(CollectionTags)

	tagIDs := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagIDs = append(tagIDs, tag.Name)
	}

	update := bson.D{{"$set", bson.M{"updatedAt": now}}, {"$push", bson.M{"recipesIDs": recipeID.Value.String()}}}
	filter := bson.M{"_id": bson.M{"$in": tagIDs}}
	_, err := tagColl.UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("recipeColl.UpdateByID: %w", err)
	}
	return nil
}
