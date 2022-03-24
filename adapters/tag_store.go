package adapters

import (
	"context"
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

	update := bson.M{"$push": bson.M{"tags": RecipeTag{
		Name: t.Name,
	}}}

	_, errOrNil := recipeColl.UpdateByID(ctx, recipeID.ID, update)
	if errOrNil != nil {
		return errOrNil
	}

	tagColl := m.ForCollection(CollectionTags)

	update = bson.M{"$push": bson.M{"recipeIds": recipeID.ID}}
	_, errOrNil = tagColl.UpdateByID(ctx, t.Name, update)
	return errOrNil
}

func (m *MongoStorage) RemoveTagFromRecipe(ctx context.Context, recipeID c.ID[c.Recipe], tagID c.ID[c.Tag]) error {
	recipeColl := m.ForCollection(CollectionRecipes)

	removeTagFromRecipe := bson.M{"$pull": bson.M{"tags": bson.M{"_id": tagID.ID}}}
	_, errOrNil := recipeColl.UpdateByID(ctx, recipeID.ID, removeTagFromRecipe)
	if errOrNil != nil {
		return errOrNil
	}

	tagsColl := m.ForCollection(CollectionRecipes)
	removeRecipeID := bson.M{"$pull": bson.M{"recipeIds": recipeID.ID}}
	_, errOrNil = tagsColl.UpdateByID(ctx, tagID.ID, removeRecipeID)
	return errOrNil
}
