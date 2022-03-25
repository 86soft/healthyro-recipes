package adapters

import (
	"context"
	d "github.com/86soft/healthyro-recipes/core"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func mapToResources(createdAt time.Time, from []d.Resource, to []Resource) {
	for _, res := range from {
		to = append(to, Resource{
			Document: Document{CreatedAt: createdAt},
			ID:       res.ID.ID,
			Name:     res.Name,
			Kind:     res.Kind,
			Value:    res.Value,
		})
	}
}

func mapFromResources(from []Resource, to []d.Resource) {
	for _, res := range from {
		to = append(to, d.Resource{
			ID:    d.FromStringID[d.Resource](res.ID),
			Name:  res.Name,
			Kind:  res.Kind,
			Value: res.Value,
		})
	}
}

func mapToRecipeTags(from []d.Tag, to []RecipeTag) {
	for _, t := range from {
		to = append(to, RecipeTag{
			Name: t.Name,
		})
	}
}

func mapFromRecipeTags(id d.ID[d.Recipe], from []RecipeTag, to []d.Tag) {
	for _, t := range from {
		to = append(to, d.Tag{
			RecipeId: id,
			Name:     t.Name,
		})
	}
}

func mapToTags(createdAt time.Time, recipe *d.Recipe, from []d.Tag, to []any) {
	for _, tag := range from {
		to = append(to, Tag{
			Document: Document{CreatedAt: createdAt},
			Name:     tag.Name,
			RecipeIDS: []string{
				recipe.ID.ID,
			},
		})
	}
}

func mapFromRecipe(from Recipe) d.Recipe {
	id := d.FromStringID[d.Recipe](from.ID)
	outRecipe := d.Recipe{
		ID:          id,
		Title:       from.Title,
		Description: from.Description,
		Resources:   make([]d.Resource, 0, len(from.Resources)),
		Tags:        make([]d.Tag, 0, len(from.Tags)),
	}
	mapFromResources(from.Resources, outRecipe.Resources)
	mapFromRecipeTags(id, from.Tags, outRecipe.Tags)
	return outRecipe
}

func mapFromRecipes(cursor *mongo.Cursor, ctx context.Context) ([]d.Recipe, error) {
	var recipes []d.Recipe
	for cursor.Next(ctx) {
		dbRecipe := Recipe{}
		err := cursor.Decode(&dbRecipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, mapFromRecipe(dbRecipe))
	}
	return recipes, nil
}
