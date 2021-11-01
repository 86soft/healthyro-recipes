package ports

import (
	"context"
	"fmt"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/domain"
	"github.com/86soft/healthyro/common"
	pb "github.com/86soft/healthyro/recipe"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}
func (s GrpcServer) ListRecipes(ctx context.Context, req *pb.ListRecipesRequest) (*pb.ListRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (s GrpcServer) GetRecipe(ctx context.Context, req *pb.GetRecipeRequest) (*pb.GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (s GrpcServer) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	cmd := command.NewCreateRecipe(req.GetTitle(), req.GetDescription(), req.GetExternalLink())
	id, err := s.app.Commands.CreateRecipe.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occured while creating new recipe")
	}
	return &pb.CreateRecipeResponse{Uuid: id.GetID().String()}, nil
}
func (s GrpcServer) UpdateRecipeTitle(ctx context.Context, req *pb.UpdateRecipeTitleRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.GetUuid())
	if err != nil {
		msg := fmt.Errorf("invalid uuid: %s", req.GetUuid())
		return nil, status.Errorf(codes.InvalidArgument, msg.Error())
	}
	rid := domain.NewRID(id)
	cmd := command.NewUpdateRecipeTitle(rid, req.GetTitle())
	err = s.app.Commands.UpdateRecipeTitle.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occured while updating recipe title")
	}
	return &common.Empty{}, nil
}
func (s GrpcServer) UpdateRecipeDescription(ctx context.Context, req *pb.UpdateRecipeDescriptionRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.GetUuid())
	if err != nil {
		msg := fmt.Errorf("invalid uuid: %s", req.GetUuid())
		return nil, status.Errorf(codes.InvalidArgument, msg.Error())
	}
	rid := domain.NewRID(id)
	cmd := command.NewUpdateRecipeDescription(rid, req.GetDescription())
	err = s.app.Commands.UpdateRecipeDescription.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occured while updating recipe description")
	}
	return &common.Empty{}, nil
}
func (s GrpcServer) UpdateRecipeExternalLink(ctx context.Context, req *pb.UpdateRecipeExternalLinkRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.GetUuid())
	if err != nil {
		msg := fmt.Errorf("invalid uuid: %s", req.GetUuid())
		return nil, status.Errorf(codes.InvalidArgument, msg.Error())
	}
	rid := domain.NewRID(id)
	cmd := command.NewUpdateRecipeExternalLink(rid, req.GetExternalLink())
	err = s.app.Commands.UpdateRecipeExternalLink.Handle(ctx, cmd)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occured while updating recipe description")
	}
	return &common.Empty{}, nil
}
func (s GrpcServer) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
