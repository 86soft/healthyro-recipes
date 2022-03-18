package adapters

import (
	"context"
	d "github.com/86soft/healthyro-recipes/domain"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *MongoStorage) AddRecipeTag(ctx context.Context, id d.ID[Recipe], t *Tag) error {
	createdAt := time.Now().UTC()
	recipeColl := m.ForCollection(CollectionRecipes)

	update := bson.M{"$push": bson.M{"tags": RecipeTag{
		ID:   t.ID,
		Name: t.Name,
	}}}

	_, errOrNil := recipeColl.UpdateByID(ctx, id.ID, update)
	if errOrNil != nil {
		return errOrNil
	}

	tagColl := m.ForCollection(CollectionTags)

	_, errOrNil = tagColl.InsertOne(ctx, Tag{
		Document: Document{CreatedAt: createdAt},
		ID:       t.ID,
		Name:     t.Name,
		RecipeIDS: []string{
			id.ID,
		},
	})
	return errOrNil
}
