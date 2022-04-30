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
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"
)

const (
	ServiceName    = "healthyro-recipes"
	ServiceVersion = "v0.0.1"
)

type Empty struct{}

var (
	_dbURL           = flag.String("conn", "", "database connection string")
	_serviceURL      = flag.String("url", "", "grpc server url")
	_developmentMode = flag.Bool("dev", false, "flag for development features")
)

type Service struct {
	serverURL string
	dbDriver  *mongo.Client
	server    *grpc.Server
	logger    zerolog.Logger
}

func setup() (*Service, error) {
	flag.Parse()
	fetchEnvVariables()
	err := validateEnvVariables()
	if err != nil {
		return nil, fmt.Errorf("validateEnvVariables: %w", err)
	}

	svc, err := newService(
		WithMongoDBClient(*_dbURL),
		WithServerURL(*_serviceURL),
	)
	if err != nil {
		return nil, fmt.Errorf("newService: %w", err)
	}

	return svc, nil
}

func (s *Service) Run() chan error {
	result := make(chan error)
	go func() {
		defer close(result)

		s.logger.Info().Msgf("starting %s", ServiceName)
		s.logger.Info().Msgf("version %s", ServiceVersion)

		if *_developmentMode {
			s.logger.Info().Msgf("~~DEVELOPMENT MODE~~")
		}

		s.logger.Info().Msgf("url: %s", s.serverURL)

		lis, err := net.Listen("tcp", fmt.Sprintf(s.serverURL))
		if err != nil {
			result <- fmt.Errorf("listen: %w", err)
			return
		}
		s.logger.Info().Msg("working...")
		result <- s.server.Serve(lis)
		return
	}()
	return result
}

func newService(args ...func(*Service) error) (*Service, error) {
	var err error
	svc := Service{}

	for _, arg := range args {
		if err := arg(&svc); err != nil {
			return nil, fmt.Errorf("args: %w", err)
		}
	}

	svc.logger = zerolog.New(os.Stdout)
	store := adapters.NewMongoStorage(svc.dbDriver)

	newApp, err := app.NewApplication(store, svc.logger)
	if err != nil {
		svc.logger.Err(err)
		return nil, err
	}

	a := ports.NewRecipeServer(newApp)

	var server *grpc.Server
	server = grpc.NewServer()

	if *_developmentMode {
		server = grpc.NewServer(grpc.UnaryInterceptor(UnaryLoggingInterceptor(svc.logger)))
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	reflection.Register(server)
	hproto.RegisterRecipeSvcServer(server, a)
	svc.server = server
	return &svc, nil
}

func (s *Service) Stop() error {
	s.logger.Warn().Msgf("%s graceful shutdown", ServiceName)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := s.dbDriver.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("dbDriver: %w", err)
	}

	stopSignal := make(chan Empty)
	go func() {
		s.server.GracefulStop()
		stopSignal <- Empty{}
	}()

	select {
	case <-stopSignal:
		return nil
	case <-time.After(15 * time.Second):
		return errors.New("GracefulStop took more than 15 seconds, timeout")
	}
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

func WithMongoDBClient(url string) func(*Service) error {
	return func(svc *Service) error {
		if url == "" {
			return errors.New("db url is empty")
		}
		var err error
		svc.dbDriver, err = adapters.NewMongoClient(url, 20)
		return err
	}
}

func WithServerURL(url string) func(*Service) error {
	return func(svc *Service) error {
		if url == "" {
			return errors.New("server url is empty")
		}
		var err error
		svc.serverURL = url
		return err
	}
}

func fetchEnvVariables() {
	if *_serviceURL == "" {
		*_serviceURL = os.Getenv("DB_URL")
	}

	if *_dbURL == "" {
		*_dbURL = os.Getenv("URL")
	}
}

func validateEnvVariables() error {
	if *_serviceURL == "" {
		return errors.New("_serviceURL is empty")
	}

	if *_dbURL == "" {
		return errors.New("_dbURL is empty")
	}
	return nil
}
