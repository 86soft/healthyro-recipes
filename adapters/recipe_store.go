package adapters

import (
	"context"
	"fmt"
	core "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (m *MongoStorage) CreateRecipe(ctx context.Context, recipe *core.Recipe) error {
	createdAt := time.Now().UTC()

	dbRecipeResources := make([]Resource, len(recipe.Resources))
	mapToResources(createdAt, recipe.Resources, dbRecipeResources)

	dbRecipeTags := make([]string, len(recipe.Tags))
	mapToRecipeTags(recipe.Tags, dbRecipeTags)

	dbRecipe := Recipe{
		Document: Document{
			CreatedAt: createdAt,
		},
		ID:          recipe.ID.Value.String(),
		Title:       recipe.Title,
		Description: recipe.Description,
		Resources:   dbRecipeResources,
		Tags:        dbRecipeTags,
	}

	recipeColl := m.ForCollection(CollectionRecipes)
	tagColl := m.ForCollection(CollectionTags)

	inRelatedTags := bson.M{"_id": dbRecipe.Tags}
	addRecipeID := bson.M{"$push": bson.M{"recipesIDs": dbRecipe.ID}}

	update := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := recipeColl.InsertOne(ctx, dbRecipe)
		if err != nil {
			return nil, fmt.Errorf("recipes: InsertOne: %w", err)
		}

		_, err = tagColl.UpdateMany(ctx, inRelatedTags, addRecipeID)
		if err != nil {
			return nil, fmt.Errorf("tag: UpdateMany: %w", err)
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

func (m *MongoStorage) AddRecipeResource(ctx context.Context, id core.ID[core.Recipe], r *core.Resource) error {
	createdAt := time.Now().UTC()
	recipeColl := m.ForCollection(CollectionRecipes)

	update := bson.M{"$push": bson.M{"resources": Resource{
		Document: Document{CreatedAt: createdAt},
		ID:       r.ID.Value.String(),
		Name:     r.Name,
		Kind:     r.Kind,
		Value:    r.Value,
	}}}
	_, errOrNil := recipeColl.UpdateByID(ctx, id.Value.String(), update)
	if errOrNil != nil {
		return errOrNil
	}
	return errOrNil
}

func (m *MongoStorage) GetRecipe(ctx context.Context, id core.ID[core.Recipe]) (core.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)

	dbRecipe := Recipe{}

	err := c.FindOne(ctx, bson.D{{"_id", id.Value.String()}}).Decode(&dbRecipe)
	if err != nil {
		return core.Recipe{}, err
	}
	recipeResources := make([]core.Resource, len(dbRecipe.Resources))

	err = mapFromResources(dbRecipe.Resources, recipeResources)
	if err != nil {
		return core.Recipe{}, fmt.Errorf("mapFromResources: %w", err)
	}

	rTags := make([]core.Tag, len(dbRecipe.Tags))
	mapFromRecipeTags(id, dbRecipe.Tags, rTags)

	return core.Recipe{
		ID:          id,
		Title:       dbRecipe.Title,
		Description: dbRecipe.Description,
		Resources:   recipeResources,
		Tags:        rTags,
	}, nil
}

func (m *MongoStorage) FindRecipesByName(ctx context.Context, title string) ([]core.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)
	cursor, err := c.Find(ctx,
		bson.D{
			{"title", bson.D{
				{"$regex", fmt.Sprintf(".*%s.*", title)},
				{"$options", "i"}, // i for case insensitive https://www.mongodb.com/docs/manual/reference/operator/query/regex/#mongodb-query-op.-regex
			}},
		})
	if err != nil {
		return nil, err
	}

	return mapFromRecipes(cursor, ctx)
}

func (m *MongoStorage) FindRecipesByTags(ctx context.Context, tags []core.Tag) ([]core.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)
	names := make([]string, 0, len(tags))
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	cursor, errOrNil := c.Aggregate(ctx, createFindRecipesByTagsPipeline(names))
	if errOrNil != nil {
		return nil, errOrNil
	}
	return mapFromRecipes(cursor, ctx)
}

func (m *MongoStorage) ListRecipes(ctx context.Context) ([]core.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)
	cursor, err := c.Find(ctx, bson.D{{}}) // should consider count, but we need pagination in future anyway
	if err != nil {
		return nil, err
	}

	return mapFromRecipes(cursor, ctx)
}

func (m *MongoStorage) UpdateRecipeTitle(ctx context.Context, id core.ID[core.Recipe], title string) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.D{{"$set", bson.D{{"title", title}}}}
	res, err := c.UpdateByID(ctx, id.Value.String(), update)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("recipe id: %s - ModifiedCount is %v, expected 1", id.Value.String(), res.ModifiedCount)
	}

	return nil
}

func (m *MongoStorage) UpdateRecipeDescription(ctx context.Context, id core.ID[core.Recipe], description string) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.D{{"$set", bson.D{{"description", description}}}}
	res, err := c.UpdateByID(ctx, id.Value.String(), update)
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("recipe id: %s - ModifiedCount is %v, expected 1", id.Value.String(), res.ModifiedCount)
	}

	return nil
}

func (m *MongoStorage) DeleteRecipe(ctx context.Context, r *core.Recipe) error {
	recipeColl := m.ForCollection(CollectionRecipes)
	tagColl := m.ForCollection(CollectionTags)

	recipeID := r.ID.Value.String()

	updateTagColl := bson.M{"$pull": bson.M{"recipesIDs": recipeID}}
	tags := make([]string, len(r.Tags))
	for i := range r.Tags {
		tags[i] = r.Tags[i].Name
	}

	filter := bson.M{"_id": bson.M{"$in": tags}}

	update := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := recipeColl.DeleteOne(ctx, bson.M{"_id": recipeID})
		if err != nil {
			return nil, fmt.Errorf("recipe: DeleteOne: %w", err)
		}

		_, err = tagColl.UpdateMany(ctx, filter, updateTagColl)
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

func (m *MongoStorage) RemoveResourceFromRecipe(ctx context.Context, recipeID core.ID[core.Recipe], resourceID core.ID[core.Resource]) error {
	c := m.ForCollection(CollectionRecipes)

	update := bson.M{"$pull": bson.M{"resources": bson.M{"_id": resourceID.Value.String()}}}
	_, errOrNil := c.UpdateByID(ctx, recipeID.Value.String(), update)
	return errOrNil
}

func (m *MongoStorage) FindRecipesByNameAndTags(ctx context.Context, name string, tags []core.Tag) ([]core.Recipe, error) {
	c := m.ForCollection(CollectionRecipes)
	tagNames := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	cursor, errOrNil := c.Aggregate(ctx, createFindRecipesByNameAndTagsPipeline(name, tagNames))
	if errOrNil != nil {
		return nil, errOrNil
	}
	return mapFromRecipes(cursor, ctx)
}

func createFindRecipesByNameAndTagsPipeline(title string, tagNames []string) mongo.Pipeline {
	filterByTags := bson.M{"tags": bson.M{"$in": tagNames}}

	filterByName := bson.D{
		{"title", bson.D{
			{"$regex", fmt.Sprintf(".*%s.*", title)},
			{"$options", "i"}, // i for case insensitive https://www.mongodb.com/docs/manual/reference/operator/query/regex/#mongodb-query-op.-regex
		}},
	}
	filters := bson.M{"$and": bson.A{filterByName, filterByTags}}

	extractSharedTagsCount := bson.D{{"$size", bson.M{"$setIntersection": bson.A{tagNames, "$tags"}}}}
	addCountToDocument := bson.M{"matchedTagCount": extractSharedTagsCount}

	sortByTagsRelevanceCount := bson.M{"matchedTagCount": -1}
	removeTemporalCount := "matchedTagCount"
	return mongo.Pipeline{match(filters), addFields(addCountToDocument), sort(sortByTagsRelevanceCount), unset(removeTemporalCount)}
}

func createFindRecipesByTagsPipeline(tagNames []string) mongo.Pipeline {
	filterByTags := bson.D{{"tags", bson.M{"$in": tagNames}}}

	extractSharedTagsCount := bson.D{{"$size", bson.M{"$setIntersection": bson.A{tagNames, "$tags"}}}}
	addCountToDocument := bson.M{"matchedTagCount": extractSharedTagsCount}

	sortByTagsRelevanceCount := bson.M{"matchedTagCount": -1}
	removeTemporalCount := "matchedTagCount"
	return mongo.Pipeline{match(filterByTags), addFields(addCountToDocument), sort(sortByTagsRelevanceCount), unset(removeTemporalCount)}
}

func match(input any) bson.D {
	return bson.D{{"$match", input}}
}

func addFields(input any) bson.D {
	return bson.D{{"$addFields", input}}
}

func sort(input any) bson.D {
	return bson.D{{"$sort", input}}
}
func unset(input any) bson.D {
	return bson.D{{"$unset", input}}
}
