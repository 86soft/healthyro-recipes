package adapters

import (
	"context"
	"fmt"
	d "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func mapToResources(createdAt time.Time, from []d.Resource, to []Resource) {
	for i, res := range from {
		to[i] = Resource{
			Document: Document{CreatedAt: createdAt},
			ID:       res.ID.Value.String(),
			Name:     res.Name,
			Kind:     res.Kind,
			Value:    res.Value,
		}
	}
}

func mapFromResources(from []Resource, to []d.Resource) error {
	for i, res := range from {
		id, err := d.FromStringID[d.Resource](res.ID)
		if err != nil {
			return fmt.Errorf("FromStringID: %w", err)
		}
		to[i] = d.Resource{
			ID:    id,
			Name:  res.Name,
			Kind:  res.Kind,
			Value: res.Value,
		}
	}
	return nil
}

func mapToRecipeTags(from []d.Tag, to []string) {
	for i, t := range from {
		to[i] = t.Name
	}
}

func mapFromRecipeTags(id d.ID[d.Recipe], from []string, to []d.Tag) {
	for i, t := range from {
		to[i] = d.Tag{
			Name: t,
		}
	}
}

func mapToTags(createdAt time.Time, recipe *d.Recipe, from []d.Tag, to []any) {
	for i, tag := range from {
		to[i] = Tag{
			Document: Document{CreatedAt: createdAt},
			Name:     tag.Name,
			RecipesIDs: []string{
				recipe.ID.Value.String(),
			},
		}
	}
}

func mapFromRecipe(from Recipe) (d.Recipe, error) {
	id, err := d.FromStringID[d.Recipe](from.ID)
	if err != nil {
		return d.Recipe{}, fmt.Errorf("FromStringID: %w", err)
	}

	outRecipe := d.Recipe{
		ID:          id,
		Title:       from.Title,
		Description: from.Description,
		Resources:   make([]d.Resource, len(from.Resources)),
		Tags:        make([]d.Tag, len(from.Tags)),
	}
	err = mapFromResources(from.Resources, outRecipe.Resources)
	if err != nil {
		return d.Recipe{}, fmt.Errorf("mapFromResources: %w", err)
	}

	mapFromRecipeTags(id, from.Tags, outRecipe.Tags)
	return outRecipe, nil
}

func mapFromRecipes(cursor *mongo.Cursor, ctx context.Context) ([]d.Recipe, error) {
	var recipes []d.Recipe
	for cursor.Next(ctx) {
		dbRecipe := Recipe{}
		err := cursor.Decode(&dbRecipe)
		if err != nil {
			return nil, fmt.Errorf("decode: %w", err)
		}
		r, err := mapFromRecipe(dbRecipe)
		if err != nil {
			return nil, fmt.Errorf("mapFromRecipe: %w", err)
		}
		recipes = append(recipes, r)
	}
	return recipes, nil
}
