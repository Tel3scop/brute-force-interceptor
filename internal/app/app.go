package app

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/Tel3scop/brute-force-interceptor/internal/closer"
	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/interceptor"
	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
	"github.com/Tel3scop/helpers/logger"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App структура приложения с сервис-провайдером и GRPC-сервером.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp вернуть новый экземпляр приложения с зависимостями.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run запуск приложения.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if _, err := config.New(); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.InitByParams(
		a.serviceProvider.Config().Log.FileName,
		a.serviceProvider.Config().Log.Level,
		a.serviceProvider.Config().Log.MaxSize,
		a.serviceProvider.Config().Log.MaxBackups,
		a.serviceProvider.Config().Log.MaxAge,
		a.serviceProvider.Config().Log.Compress,
		a.serviceProvider.Config().Log.StdOut,
	)
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			interceptor.ValidateInterceptor,
			interceptor.LogInterceptor,
		)),
	)

	reflection.Register(a.grpcServer)
	accessAPI.RegisterAntiBruteforceServer(a.grpcServer, a.serviceProvider.AccessImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.Config().GRPC.Address)

	list, err := net.Listen("tcp", a.serviceProvider.Config().GRPC.Address)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
