package ports

import (
	"context"
	"github.com/86soft/healthyro-recipes/app"
	hproto "github.com/86soft/healthyro-recipes/ports/protos"
)

type RecipeServer struct {
	app app.Application
	hproto.UnimplementedRecipeServiceServer
}

func NewRecipeServer(application app.Application) RecipeServer {
	return RecipeServer{app: application}
}

func (r RecipeServer) AddRecipe(ctx context.Context, request *hproto.AddRecipeRequest) (*hproto.AddRecipeResponse, error) {
	/*cmd := commands.NewAddRecipe("test", "kappa")
	err := r.app.Commands.AddRecipe.Handle(ctx, cmd)*/
	return &hproto.AddRecipeResponse{}, nil
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
