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
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"
)

const (
	ServiceName    = "healthyro-recipes"
	ServiceVersion = "v0.0.1"
)

type Service struct {
	isLocal    string
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

func setup() (*Service, error) {
	// TODO: once the is a stable mvp, refactor this so it is not that ugly
	svc := createDefaultSvc()
	local := flag.Bool("local", false,
		"determines if service is running locally, if so it needs passed arguments")

	customConn := flag.String("conn", "", "adapters database connection string")
	customGrpcPort := flag.String("grpcPort", "", "grpc server port")

	flag.Parse()

	svc.log = zerolog.New(os.Stdout)
	if *local {
		svc.mongoDbUrl = *customConn
		svc.grpcPort = *customGrpcPort
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMilli}
		svc.log = zerolog.New(output).With().Timestamp().Logger()
	}

	if svc.mongoDbUrl == "" {
		return nil, errors.New("adapters connection string is missing")
	}
	if svc.grpcPort == "" {
		return nil, errors.New("grpcPort port is missing")
	}

	var err error
	svc.dbClient, err = adapters.NewMongoClient(svc.mongoDbUrl, 20)
	if err != nil {
		return nil, fmt.Errorf("adapters: %w", err)
	}

	store := adapters.NewMongoStorage(svc.dbClient)
	newApp, err := app.NewApplication(store, svc.log)
	if err != nil {
		svc.log.Err(err)
		return nil, err
	}
	server := ports.NewRecipeServer(newApp)
	svc.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(UnaryLoggingInterceptor(svc.log)))

	reflection.Register(svc.grpcServer)
	hproto.RegisterRecipeSvcServer(svc.grpcServer, server)
	mongoInfi := fmt.Sprintf("adapters connection: %s", svc.mongoDbUrl)
	grpcInfo := fmt.Sprintf("grpc listening on port: %s", svc.grpcPort)
	svc.log.Info().Msg(mongoInfi)
	svc.log.Info().Msg(grpcInfo)
	return &svc, nil
}

func createDefaultSvc() Service {
	svc := Service{}
	svc.mongoDbUrl = os.Getenv("MONGO_URL")
	svc.grpcPort = os.Getenv("PORT")
	svc.isLocal = os.Getenv("LOCAL_MODE")
	return svc
}

func (s *Service) Clear() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return s.dbClient.Disconnect(ctx)
}

func UnaryLoggingInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now().UTC()

		h, err := handler(ctx, req)
		took := time.Since(start)
		if err != nil {
			logger.Error().Msg(err.Error())
		}
		logger.Info().Msg(fmt.Sprintf("Request: %s took %v", info.FullMethod, took))
		return h, err
	}
}
