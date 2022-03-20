package adapters

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *MongoStorage) AddRecipeTag(ctx context.Context, id d.ID[d.Recipe], t *d.Tag) error {
	createdAt := time.Now().UTC()
	recipeColl := m.ForCollection(CollectionRecipes)

	update := bson.M{"$push": bson.M{"tags": RecipeTag{
		ID:   t.ID.ID,
		Name: t.Name,
	}}}

	_, errOrNil := recipeColl.UpdateByID(ctx, id.ID, update)
	if errOrNil != nil {
		return errOrNil
	}

	tagColl := m.ForCollection(CollectionTags)

	_, errOrNil = tagColl.InsertOne(ctx, Tag{
		Document: Document{CreatedAt: createdAt},
		ID:       t.ID.ID,
		Name:     t.Name,
		RecipeIDS: []string{
			id.ID,
		},
	})
	return errOrNil
}

func (m *MongoStorage) DeleteRecipeTag(ctx context.Context, recipeID d.ID[d.Recipe], tagID d.ID[d.Tag]) error {
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
