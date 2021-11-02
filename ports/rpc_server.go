package ports

import (
	"context"
	"github.com/86soft/healthyro-recipes/Helpers"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/app/query"
	"github.com/86soft/healthyro/common"
	pb "github.com/86soft/healthyro/recipe"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecipeServer struct {
	app app.Application
	pb.UnimplementedRecipeServiceServer
}

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (s RecipeServer) ListRecipes(ctx context.Context, req *pb.ListRecipesRequest) (*pb.ListRecipesResponse, error) {
	recipes, err := s.app.Queries.ListRecipes.Handle(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occurred while querying all recipes")
	}
	return &pb.ListRecipesResponse{Recipes: Helpers.MapRecipesToProto(recipes)}, nil
}

func (s RecipeServer) GetRecipe(ctx context.Context, req *pb.GetRecipeRequest) (*pb.GetRecipeResponse, error) {
	q := query.NewGetRecipeById(req.GetUuid())

	recipe, err := s.app.Queries.GetRecipeById.Handle(ctx, q)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occurred while querying new recipe")
	}
	return &pb.GetRecipeResponse{Recipe: Helpers.MapRecipeToProto(&recipe)}, nil
}

func (s RecipeServer) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	cmd := command.NewCreateRecipe(req.GetTitle(), req.GetDescription(), req.GetExternalLink())

	id, err := s.app.Commands.CreateRecipe.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occurred while creating new recipe")
	}

	return &pb.CreateRecipeResponse{Uuid: id.GetID().String()}, nil
}

func (s RecipeServer) UpdateRecipeTitle(ctx context.Context, req *pb.UpdateRecipeTitleRequest) (*common.Empty, error) {
	cmd := command.NewUpdateRecipeTitle(req.GetUuid(), req.GetTitle())

	err := s.app.Commands.UpdateRecipeTitle.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while updating recipe title")
	}

	return &common.Empty{}, nil
}

func (s RecipeServer) UpdateRecipeDescription(ctx context.Context, req *pb.UpdateRecipeDescriptionRequest) (*common.Empty, error) {
	cmd := command.NewUpdateRecipeDescription(req.GetUuid(), req.GetDescription())

	err := s.app.Commands.UpdateRecipeDescription.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while updating recipe description")
	}

	return &common.Empty{}, nil
}

func (s RecipeServer) UpdateRecipeExternalLink(ctx context.Context, req *pb.UpdateRecipeExternalLinkRequest) (*common.Empty, error) {
	cmd := command.NewUpdateRecipeExternalLink(req.GetUuid(), req.GetExternalLink())

	err := s.app.Commands.UpdateRecipeExternalLink.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while updating recipe description")
	}

	return &common.Empty{}, nil
}

func (s RecipeServer) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*common.Empty, error) {
	cmd := command.NewDeleteRecipe(req.GetUuid())

	err := s.app.Commands.DeleteRecipe.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while deleting recipe")
	}
	return &common.Empty{}, nil
}

/*func RunGRPCServer(registerServer func(server *grpc.Server)) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)
	listen, err := net.Listen("tcp", addr)
}*/
