package app

import (
	"context"
	"filters/app/internal/config"
	"filters/app/internal/domain/service/filter_service"
	"fmt"
	"net"
	"strconv"

	"filters/app/internal/data_provider/filter_data_provider"

	"time"

	pb "filters/app/gen/proto"

	"github.com/marketconnect/db_client/postgresql"
	"github.com/marketconnect/jwt_manager"
	"github.com/marketconnect/logger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

type App struct {
	cfg        *config.Config
	logger     logger.Logger
	grpcServer *grpc.Server
}

func NewApp(config *config.Config, logger logger.Logger) (App, error) {
	logger.Info("Postgres initializing")
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.PostgreUsername, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}
	tokenDuration, err := strconv.Atoi(config.Jwt.TokenDuration)
	if err != nil {
		logger.Fatal(err)
	}

	jwtManager := jwt_manager.NewJWTManager(config.Jwt.SecretKey, time.Duration((time.Minute * time.Duration(tokenDuration))))

	// Data Providers

	filterDataProvider := filter_data_provider.NewFilterStorage(pgClient)

	// Services

	filterService := filter_service.NewFilterService(filterDataProvider, jwtManager, logger)

	// interceptor := auth_interceptor.NewAuthInterceptor(subscriptionDataProvider, tokenManager)

	// grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	grpcServer := grpc.NewServer()

	pb.RegisterFilterServiceServer(grpcServer, filterService)

	// Включение Reflection
	// reflection.Register(grpcServer)

	return App{
		cfg:        config,
		logger:     logger,
		grpcServer: grpcServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startGRPC(ctx)
	})
	return grp.Wait()
}

func (a *App) startGRPC(ctx context.Context) error {
	a.logger.Info("start GRPC")
	address := fmt.Sprintf("%s:%s", a.cfg.GRPC.IP, a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		a.logger.Fatal("cannot start GRPC server: ", err)
	}
	a.logger.Info("start GRPC server on address %s", address)
	err = a.grpcServer.Serve(listener)
	if err != nil {
		a.logger.Fatal("cannot start GRPC server: ", err)
	}
	return nil
}
