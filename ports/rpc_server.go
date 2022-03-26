package ports

import (
	c "context"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/commands"
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/core"
	p "github.com/86soft/healthyro-recipes/ports/protos"
)

type RecipeServer struct {
	p.UnimplementedRecipeSvcServer
	app app.Application
}

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (r RecipeServer) CreateRecipe(ctx c.Context, req *p.CreateRecipeRequest) (*p.CreateRecipeResponse, error) {
	rpcRes := req.GetResources()
	res := make([]commands.RecipeResources, 0, len(rpcRes))
	for _, r := range rpcRes {
		res = append(res, commands.RecipeResources{
			Name:  r.GetName(),
			Kind:  r.GetKind(),
			Value: r.GetValue(),
		})
	}

	rpcTags := req.GetTags()
	tags := make([]string, 0, len(rpcTags))
	for _, t := range rpcTags {
		tags = append(tags, t.GetName())
	}

	cmd := commands.CreateRecipe{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Resources:   res,
		Tags:        tags,
	}
	id, err := r.app.Commands.CreateRecipe.Handle(ctx, cmd)
	if err != nil {
		r.app.Log.Error().Msg(err.Error())
		return nil, err
	}

	return &p.CreateRecipeResponse{
		RecipeId: id.ID,
	}, nil
}

func (r RecipeServer) ListRecipe(ctx c.Context, _ *p.ListRecipeRequest) (*p.ListRecipeResponse, error) {
	recipes, err := r.app.Queries.ListRecipes.Handle(ctx)
	if err != nil {
		return nil, err
	}
	rspRecipes := mapRecipesToResponse(recipes)
	return &p.ListRecipeResponse{Recipes: rspRecipes}, nil
}

func (r RecipeServer) FindRecipesByName(ctx c.Context, req *p.FindRecipesByNameRequest) (*p.FindRecipesByNameResponse, error) {
	cmd := queries.FindRecipesByName{Name: req.GetName()}
	recipes, err := r.app.Queries.FindRecipesByName.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}
	rsp := mapRecipesToResponse(recipes)
	return &p.FindRecipesByNameResponse{Recipes: rsp}, nil
}

func (r RecipeServer) GetRecipe(ctx c.Context, req *p.GetRecipeRequest) (*p.GetRecipeResponse, error) {
	q := queries.GetRecipeById{
		RecipeID: core.FromStringID[core.Recipe](req.GetRecipeId()),
	}
	recipe, err := r.app.Queries.GetRecipeById.Handle(ctx, q)
	if err != nil {
		return nil, err
	}
	return &p.GetRecipeResponse{Recipe: mapRecipeToResponse(recipe)}, nil

}

func (r RecipeServer) FindRecipesByTags(ctx c.Context, req *p.FindRecipesByTagsRequest) (*p.FindRecipesByTagsResponse, error) {
	reqTags := req.GetTags()
	t := make([]core.Tag, 0, len(reqTags))
	for _, tag := range reqTags {
		t = append(t, core.Tag{
			Name: tag,
		})
	}
	q := queries.FindRecipesByTags{Tags: t}
	recipes, err := r.app.Queries.FindRecipesByTags.Handle(ctx, q)
	if err != nil {
		return nil, err
	}
	return &p.FindRecipesByTagsResponse{Recipes: mapRecipesToResponse(recipes)}, nil
}

func (r RecipeServer) UpdateRecipeTitle(ctx c.Context, req *p.UpdateRecipeTitleRequest) (*p.UpdateRecipeTitleResponse, error) {
	cmd := commands.UpdateRecipeTitle{
		RecipeID: core.FromStringID[core.Recipe](req.GetRecipeId()),
		Title:    req.GetTitle(),
	}
	return &p.UpdateRecipeTitleResponse{}, r.app.Commands.UpdateRecipeTitle.Handle(ctx, cmd)
}

func (r RecipeServer) UpdateRecipeDescription(ctx c.Context, req *p.UpdateRecipeDescriptionRequest) (*p.UpdateRecipeDescriptionResponse, error) {
	cmd := commands.UpdateRecipeDescription{
		RecipeID:    core.FromStringID[core.Recipe](req.GetRecipeId()),
		Description: req.GetDescription(),
	}
	return &p.UpdateRecipeDescriptionResponse{}, r.app.Commands.UpdateRecipeDescription.Handle(ctx, cmd)
}

func (r RecipeServer) DeleteRecipe(ctx c.Context, req *p.DeleteRecipeRequest) (*p.DeleteRecipeResponse, error) {
	cmd := commands.DeleteRecipe{
		RecipeID: core.FromStringID[core.Recipe](req.GetRecipeId()),
	}
	return &p.DeleteRecipeResponse{}, r.app.Commands.DeleteRecipe.Handle(ctx, cmd)
}

func (r RecipeServer) RemoveRecipeFromResource(ctx c.Context, req *p.RemoveResourceFromRecipeRequest) (*p.RemoveRecipeFromResourceResponse, error) {
	cmd := commands.RemoveResourceFromRecipe{
		RecipeID:   core.FromStringID[core.Recipe](req.GetRecipeId()),
		ResourceID: core.FromStringID[core.Resource](req.GetResourceId()),
	}
	return &p.RemoveRecipeFromResourceResponse{}, r.app.Commands.RemoveResourceFromRecipe.Handle(ctx, cmd)
}

func (r RecipeServer) AddRecipeResource(ctx c.Context, req *p.AddRecipeResourceRequest) (*p.AddRecipeResourceResponse, error) {
	cmd := commands.AddRecipeResource{
		Name:     req.GetName(),
		Kind:     req.GetKind(),
		Value:    req.GetValue(),
		RecipeID: core.FromStringID[core.Recipe](req.GetRecipeId()),
	}
	return &p.AddRecipeResourceResponse{}, r.app.Commands.AddRecipeResource.Handle(ctx, cmd)
}

func (r RecipeServer) CreateTag(ctx c.Context, req *p.CreateTagRequest) (*p.CreateTagResponse, error) {
	cmd := commands.CreateTag{Name: req.GetName()}
	id, err := r.app.Commands.CreateTag.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return &p.CreateTagResponse{TagId: id.ID}, nil
}
func (r RecipeServer) AddTagToRecipe(ctx c.Context, req *p.AddTagToRecipeRequest) (*p.AddTagToRecipeResponse, error) {
	cmd := commands.AddTagToRecipe{
		Name:         req.GetTagName(),
		TagID:        core.FromStringID[core.Tag](req.GetTagId()),
		RecipeID:     core.FromStringID[core.Recipe](req.GetRecipeId()),
		CreateNewTag: req.GetCreateNewTag(),
	}
	return &p.AddTagToRecipeResponse{}, r.app.Commands.AddTagToRecipe.Handle(ctx, cmd)
}
func (r RecipeServer) RemoveTagFromRecipe(ctx c.Context, req *p.RemoveTagFromRecipeRequest) (*p.RemoveTagFromRecipeResponse, error) {
	cmd := commands.RemoveTagFromRecipe{
		TagID:    core.FromStringID[core.Tag](req.GetTagId()),
		RecipeID: core.FromStringID[core.Recipe](req.GetRecipeId()),
	}
	return &p.RemoveTagFromRecipeResponse{}, r.app.Commands.RemoveTagFromRecipe.Handle(ctx, cmd)
}

func mapRecipesToResponse(recipes []core.Recipe) []*p.Recipe {
	rspRecipes := make([]*p.Recipe, 0, len(recipes))
	for _, r := range recipes {
		rspRecipes = append(rspRecipes, mapRecipeToResponse(r))
	}
	return rspRecipes
}
func mapRecipeToResponse(r core.Recipe) *p.Recipe {
	resources := make([]*p.Recipe_Resource, 0, len(r.Resources))
	for _, res := range r.Resources {
		resources = append(resources, &p.Recipe_Resource{
			Name:  res.Name,
			Kind:  res.Kind,
			Value: res.Value,
		})
	}
	tags := make([]*p.Recipe_Tag, 0, len(r.Tags))
	for _, res := range r.Tags {
		tags = append(tags, &p.Recipe_Tag{
			Name: res.Name,
		})
	}
	return &p.Recipe{
		RecipeId:    r.ID.ID,
		Title:       r.Title,
		Description: r.Description,
		Resources:   resources,
		Tags:        tags,
	}
}
