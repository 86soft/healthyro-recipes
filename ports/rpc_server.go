package ports

import (
	"context"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/commands"
	hproto "github.com/86soft/healthyro-recipes/ports/protos"
)

type RecipeServer struct {
	app app.Application
	hproto.UnimplementedRecipeServer
}

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (r RecipeServer) AddRecipe(ctx context.Context, req *hproto.AddRecipeRequest) (*hproto.AddRecipeResponse, error) {
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

	cmd := commands.AddRecipe{
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

	return &hproto.AddRecipeResponse{
		RecipeId: id.GetID(),
	}, nil
}
