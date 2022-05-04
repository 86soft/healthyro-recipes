package adapters

import (
	"context"
	"fmt"
	c "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	tagColl := m.ForCollection(CollectionTags)

	updateRecipe := bson.M{"$push": bson.M{"tags": t.Name}}
	updateTagColl := bson.M{"$push": bson.M{"recipesIDs": recipeID.Value.String()}}

	update := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := recipeColl.UpdateByID(sessCtx, recipeID.Value.String(), updateRecipe)
		if err != nil {
			return nil, fmt.Errorf("recipe: UpdateByID: %w", err)
		}

		_, err = tagColl.UpdateByID(ctx, t.Name, updateTagColl)
		if err != nil {
			return nil, fmt.Errorf("tag: UpdateByID: %w", err)
		}
		return nil, nil
	}

	session, err := m.client.StartSession()
	if err != nil {
		return fmt.Errorf("StartSession: %w", err)
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, update)
	if err != nil {
		return fmt.Errorf("WithTransaction: %w", err)
	}

	return nil
}

func (m *MongoStorage) RemoveTagFromRecipe(ctx context.Context, recipeID c.ID[c.Recipe], tagName string) error {
	recipeColl := m.ForCollection(CollectionRecipes)
	tagColl := m.ForCollection(CollectionTags)

	update := func(sessCtx mongo.SessionContext) (interface{}, error) {
		updateRecipe := bson.M{"$pull": bson.M{"tags": tagName}}
		_, err := recipeColl.UpdateByID(sessCtx, recipeID.Value.String(), updateRecipe)
		if err != nil {
			return nil, fmt.Errorf("recipe: UpdateByID: %w", err)
		}

		updateTagColl := bson.M{"$pull": bson.M{"recipesIDs": recipeID.Value.String()}}
		_, err = tagColl.UpdateByID(ctx, tagName, updateTagColl)
		if err != nil {
			return nil, fmt.Errorf("tag: UpdateByID: %w", err)
		}
		return nil, nil
	}

	session, err := m.client.StartSession()
	if err != nil {
		return fmt.Errorf("StartSession: %w", err)
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, update)
	if err != nil {
		return fmt.Errorf("WithTransaction: %w", err)
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

func (m *MongoStorage) CheckRecipeForTag(ctx context.Context, recipeID c.ID[c.Recipe], tagName string) (bool, error) {
	tagColl := m.ForCollection(CollectionTags)
	filter := bson.M{"_id": tagName, "recipesIDs": bson.M{"$all": bson.A{recipeID.Value.String()}}}

	count, err := tagColl.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("CountDocuments: %w", err)
	}

	if count == 0 {
		return false, nil
	}
	if count > 1 {
		return false, fmt.Errorf(
			"invalid tag state, tag:%s | recipeID: %s | recipeID count: %v",
			tagName,
			recipeID.Value.String(),
			count,
		)
	}
	return true, nil
}

func (m *MongoStorage) TagExist(ctx context.Context, tagName string) (bool, error) {
	tagColl := m.ForCollection(CollectionTags)
	filter := bson.M{"_id": tagName}
	count, err := tagColl.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("CountDocuments: %w", err)
	}

	return count == 1, nil
}
