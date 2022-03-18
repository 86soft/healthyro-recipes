package adapters

import (
	d "github.com/86soft/healthyro-recipes/domain"
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
			ID:    d.FromStringID[Resource](res.ID),
			Name:  res.Name,
			Kind:  res.Kind,
			Value: res.Value,
		})
	}
}

func mapToRecipeTags(from []d.Tag, to []RecipeTag) {
	for _, t := range from {
		to = append(to, RecipeTag{
			ID:   t.ID.ID,
			Name: t.Name,
		})
	}
}

func mapFromRecipeTags(id d.ID[d.Recipe], from []RecipeTag, to []d.Tag) {
	for _, t := range from {
		to = append(to, d.Tag{
			ID:       d.FromStringID[Tag](t.ID),
			RecipeId: id,
			Name:     t.Name,
		})
	}
}

func mapToTags(createdAt time.Time, recipe *d.Recipe, from []d.Tag, to []any) {
	for _, tag := range from {
		to = append(to, Tag{
			Document: Document{CreatedAt: createdAt},
			ID:       tag.ID.ID,
			Name:     tag.Name,
			RecipeIDS: []string{
				recipe.ID.ID,
			},
		})
	}
}
