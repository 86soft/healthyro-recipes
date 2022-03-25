package ports

import (
	c "context"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/commands"
	. "github.com/86soft/healthyro-recipes/ports/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	UnimplementedRecipeSvcServer
	app app.Application
}

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (r RecipeServer) CreateRecipe(ctx c.Context, req *CreateRecipeRequest) (*CreateRecipeResponse, error) {
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
	id, err := r.app.Commands.AddRecipe.Handle(ctx, cmd)
	if err != nil {
		r.app.Log.Error().Msg(err.Error())
		return nil, err
	}

	return &CreateRecipeResponse{
		RecipeId: id.ID,
	}, nil
}

func (r RecipeServer) ListRecipe(ctx c.Context, req *ListRecipeRequest) (*ListRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipe not implemented")
}
func (r RecipeServer) FindRecipesByName(c.Context, *FindRecipesByNameRequest) (*FindRecipesByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRecipesByName not implemented")
}
func (r RecipeServer) GetRecipe(c.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (r RecipeServer) FindRecipesByTags(c.Context, *FindRecipesByTagsRequest) (*FindRecipesByTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRecipesByTags not implemented")
}
func (r RecipeServer) UpdateRecipeTitle(c.Context, *UpdateRecipeTitleRequest) (*UpdateRecipeTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipeTitle not implemented")
}
func (r RecipeServer) UpdateRecipeDescription(c.Context, *UpdateRecipeDescriptionRequest) (*UpdateRecipeDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipeDescription not implemented")
}
func (r RecipeServer) DeleteRecipe(c.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (r RecipeServer) DeleteRecipeResource(c.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveResourceFromRecipe not implemented")
}
func (r RecipeServer) AddRecipeResource(c.Context, *AddRecipeResourceRequest) (*AddRecipeResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecipeResource not implemented")
}
func (r RecipeServer) CreateTag(c.Context, *CreateTagRequest) (*CreateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (r RecipeServer) AddTagToRecipe(c.Context, *AddTagToRecipeRequest) (*AddTagToRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTagToRecipe not implemented")
}
func (r RecipeServer) RemoveTagFromRecipe(c.Context, *RemoveTagFromRecipeRequest) (*RemoveTagFromRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTagFromRecipe not implemented")
}
