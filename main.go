package main

import (
	"fmt"
	"github.com/86soft/healthyro-recipes/adapters/dao"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/ports"
	"github.com/86soft/healthyro/recipe"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func main() {

	//dsn := os.Getenv("db_conn")
	dsn := os.Getenv("db_conn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to db")
	}
	logger := zerolog.New(os.Stderr)
	a := app.NewApplication(dao.NewRecipeRepository(db))
	srv := ports.NewRecipeServer(a)
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	recipe.RegisterRecipeServiceServer(grpcServer, srv)
	logger.Fatal().Msg(grpcServer.Serve(lis).Error())
}
