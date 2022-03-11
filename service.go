package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/86soft/healthyro-recipes/adapters"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/ports"
	hproto "github.com/86soft/healthyro-recipes/ports/protos"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

const (
	ServiceName    = "healthyro-recipes"
	ServiceVersion = "v0.0.1"
)

type Service struct {
	mongoDbUrl string
	grpcPort   string
	grpcServer *grpc.Server
	dbClient   *mongo.Client
	log        zerolog.Logger
}

func (s *Service) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", s.grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	return s.grpcServer.Serve(lis)
}

func Setup(logger zerolog.Logger) (*Service, error) {
	svc := Service{log: logger}
	svc.mongoDbUrl = os.Getenv("MONGO_URL")
	svc.grpcPort = os.Getenv("PORT")

	local := flag.Bool("local", false,
		"determines if service is running locally, if so it needs passed arguments")
	customConn := flag.String("conn", "", "mongo database connection string")
	customGrpcPort := flag.String("grpcPort", "", "grpc server port")

	flag.Parse()

	if *local {
		svc.mongoDbUrl = *customConn
		svc.grpcPort = *customGrpcPort
	}
	if svc.mongoDbUrl == "" {
		return nil, errors.New("mongo connection string is missing")
	}
	if svc.grpcPort == "" {
		return nil, errors.New("grpcPort port is missing")
	}

	var err error
	svc.dbClient, err = adapters.NewMongoClient(svc.mongoDbUrl, 20)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	store := adapters.NewMongoStorage(svc.dbClient)
	newApp := app.NewApplication(store)
	server := ports.NewRecipeServer(newApp)
	svc.grpcServer = grpc.NewServer()

	hproto.RegisterRecipeServiceServer(svc.grpcServer, server)

	return &svc, nil
}

func (s *Service) Clear() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return s.dbClient.Disconnect(ctx)
}
