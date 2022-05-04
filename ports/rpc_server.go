package ports

import (
	c "context"
	"fmt"
	"github.com/86soft/healthyro-recipes/app"
	cmnds "github.com/86soft/healthyro-recipes/app/commands"
	"github.com/86soft/healthyro-recipes/app/queries"
	"github.com/86soft/healthyro-recipes/core"
	p "github.com/86soft/healthyro-recipes/ports/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	p.UnimplementedRecipeSvcServer
	app app.Application
}

var _ p.RecipeSvcServer = (*RecipeServer)(nil)

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (r RecipeServer) CreateRecipe(ctx c.Context, req *p.CreateRecipeRequest) (*p.CreateRecipeResponse, error) {
	rpcRes := req.GetResources()
	res := make([]cmnds.RecipeResources, 0, len(rpcRes))
	for _, r := range rpcRes {
		res = append(res, cmnds.RecipeResources{
			Name:  r.GetName(),
			Kind:  r.GetKind(),
			Value: r.GetValue(),
		})
	}

	rpcTags := req.GetTags()
	tags := make([]cmnds.CreateRecipeCmdTag, 0, len(rpcTags))
	for _, t := range rpcTags {
		tags = append(tags, cmnds.CreateRecipeCmdTag{
			Name:   t.GetName(),
			Create: t.GetCreate(),
		})
	}

	cmd := cmnds.CreateRecipe{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Resources:   res,
		Tags:        tags,
	}
	id, err := r.app.CreateRecipe(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("CreateRecipe: %w", err)
	}

	return &p.CreateRecipeResponse{
		RecipeId: id.Value.String(),
	}, nil
}

func (r RecipeServer) ListRecipe(ctx c.Context, _ *p.ListRecipeRequest) (*p.ListRecipeResponse, error) {
	recipes, err := r.app.ListRecipes(ctx)
	if err != nil {
		return nil, err
	}
	rspRecipes := mapRecipesToResponse(recipes)
	return &p.ListRecipeResponse{Recipes: rspRecipes}, nil
}

func (r RecipeServer) FindRecipesByName(ctx c.Context, req *p.FindRecipesByNameRequest) (*p.FindRecipesByNameResponse, error) {
	cmd := queries.FindRecipesByName{Name: req.GetName()}
	recipes, err := r.app.FindRecipesByName(ctx, cmd)
	if err != nil {
		return nil, err
	}
	rsp := mapRecipesToResponse(recipes)
	return &p.FindRecipesByNameResponse{Recipes: rsp}, nil
}

func (r RecipeServer) GetRecipe(ctx c.Context, req *p.GetRecipeRequest) (*p.GetRecipeResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}
	q := queries.GetRecipeById{
		RecipeID: recipeID,
	}
	recipe, err := r.app.GetRecipeById(ctx, q)
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
	recipes, err := r.app.FindRecipesByTags(ctx, q)
	if err != nil {
		return nil, err
	}
	return &p.FindRecipesByTagsResponse{Recipes: mapRecipesToResponse(recipes)}, nil
}

func (r RecipeServer) UpdateRecipeTitle(ctx c.Context, req *p.UpdateRecipeTitleRequest) (*p.UpdateRecipeTitleResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.UpdateRecipeTitle{
		RecipeID: recipeID,
		Title:    req.GetTitle(),
	}
	return &p.UpdateRecipeTitleResponse{}, r.app.UpdateRecipeTitle(ctx, cmd)
}

func (r RecipeServer) UpdateRecipeDescription(ctx c.Context, req *p.UpdateRecipeDescriptionRequest) (*p.UpdateRecipeDescriptionResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.UpdateRecipeDescription{
		RecipeID:    recipeID,
		Description: req.GetDescription(),
	}
	return &p.UpdateRecipeDescriptionResponse{}, r.app.UpdateRecipeDescription(ctx, cmd)
}

func (r RecipeServer) DeleteRecipe(ctx c.Context, req *p.DeleteRecipeRequest) (*p.DeleteRecipeResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.DeleteRecipe{
		RecipeID: recipeID,
	}
	return &p.DeleteRecipeResponse{}, r.app.DeleteRecipe(ctx, cmd)
}

func (r RecipeServer) AddRecipeResource(ctx c.Context, req *p.AddRecipeResourceRequest) (*p.AddRecipeResourceResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.AddRecipeResource{
		Name:     req.GetName(),
		Kind:     req.GetKind(),
		Value:    req.GetValue(),
		RecipeID: recipeID,
	}
	return &p.AddRecipeResourceResponse{}, r.app.AddRecipeResource(ctx, cmd)
}

func (r RecipeServer) CreateTag(ctx c.Context, req *p.CreateTagRequest) (*p.CreateTagResponse, error) {
	cmd := cmnds.CreateTag{Name: req.GetName()}
	id, err := r.app.CreateTag(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &p.CreateTagResponse{TagId: id.Value.String()}, nil
}
func (r RecipeServer) AddTagToRecipe(ctx c.Context, req *p.AddTagToRecipeRequest) (*p.AddTagToRecipeResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.AddTagToRecipe{
		Name:         req.GetTagName(),
		RecipeID:     recipeID,
		CreateNewTag: req.GetCreateNewTag(),
	}
	return &p.AddTagToRecipeResponse{}, r.app.AddTagToRecipe(ctx, cmd)
}
func (r RecipeServer) RemoveTagFromRecipe(ctx c.Context, req *p.RemoveTagFromRecipeRequest) (*p.RemoveTagFromRecipeResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeID())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}

	cmd := cmnds.RemoveTagFromRecipe{
		Tag:      req.GetTag(),
		RecipeID: recipeID,
	}
	return &p.RemoveTagFromRecipeResponse{}, r.app.RemoveTagFromRecipe(ctx, cmd)
}

func (r RecipeServer) FindRecipesByNameAndTags(
	ctx c.Context,
	req *p.FindRecipesByNameAndTagsRequest,
) (*p.FindRecipesByNameAndTagsResponse, error) {
	reqTags := req.GetTags()
	t := make([]core.Tag, 0, len(reqTags))
	for _, tag := range reqTags {
		t = append(t, core.Tag{
			Name: tag,
		})
	}

	recipes, err := r.app.FindRecipesByNameAndTags(ctx, queries.FindRecipesByNameAndTags{
		Name: req.GetName(),
		Tags: t,
	})
	if err != nil {
		return nil, err
	}
	return &p.FindRecipesByNameAndTagsResponse{Recipes: mapRecipesToResponse(recipes)}, nil
}

func (r RecipeServer) RemoveResourceFromRecipe(
	ctx c.Context,
	req *p.RemoveResourceFromRecipeRequest,
) (*p.RemoveRecipeFromResourceResponse, error) {
	recipeID, err := core.FromStringID[core.Recipe](req.GetRecipeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize recipe id")
	}
	resID, err := core.FromStringID[core.Resource](req.GetResourceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not deserialize resource id")
	}

	cmd := cmnds.RemoveResourceFromRecipe{
		RecipeID:   recipeID,
		ResourceID: resID,
	}
	return &p.RemoveRecipeFromResourceResponse{}, r.app.RemoveResourceFromRecipe(ctx, cmd)
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
		RecipeId:    r.ID.Value.String(),
		Title:       r.Title,
		Description: r.Description,
		Resources:   resources,
		Tags:        tags,
	}
}
