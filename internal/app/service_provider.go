package app

import (
	"context"
	"log"

	"github.com/Tel3scop/brute-force-interceptor/internal/api/access"
	"github.com/Tel3scop/brute-force-interceptor/internal/client/db"
	"github.com/Tel3scop/brute-force-interceptor/internal/client/db/pg"
	"github.com/Tel3scop/brute-force-interceptor/internal/client/db/transaction"
	"github.com/Tel3scop/brute-force-interceptor/internal/closer"
	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	blackListRepo "github.com/Tel3scop/brute-force-interceptor/internal/repository/blacklist"
	bucketRepo "github.com/Tel3scop/brute-force-interceptor/internal/repository/bucket"
	whiteListRepo "github.com/Tel3scop/brute-force-interceptor/internal/repository/whitelist"
	"github.com/Tel3scop/brute-force-interceptor/internal/service"
	accessService "github.com/Tel3scop/brute-force-interceptor/internal/service/access"
	bucketService "github.com/Tel3scop/brute-force-interceptor/internal/service/bucket"
	"github.com/go-redis/redis/v8"
)

type serviceProvider struct {
	config *config.Config

	dbClient            db.Client
	txManager           db.TxManager
	whiteListRepository repository.WhiteListRepository
	blackListRepository repository.BlackListRepository
	bucketRepository    repository.BucketRepository

	accessService service.AccessService
	bucketService service.BucketService

	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := config.New()
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}
		s.config = cfg
	}
	return s.config
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.Config().Postgres.DSN)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) WhiteListRepository(ctx context.Context) repository.WhiteListRepository {
	if s.whiteListRepository == nil {
		s.whiteListRepository = whiteListRepo.NewRepository(s.DBClient(ctx))
	}

	return s.whiteListRepository
}

func (s *serviceProvider) BlackListRepository(ctx context.Context) repository.BlackListRepository {
	if s.blackListRepository == nil {
		s.blackListRepository = blackListRepo.NewRepository(s.DBClient(ctx))
	}

	return s.blackListRepository
}

func (s *serviceProvider) BucketRepository() repository.BucketRepository {
	if s.bucketRepository == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     s.Config().Redis.Address,
			Password: s.Config().Redis.Password,
			DB:       s.Config().Redis.DB,
		})

		s.bucketRepository = bucketRepo.NewRepository(client)
	}

	return s.bucketRepository
}

func (s *serviceProvider) BucketService() service.BucketService {
	if s.bucketService == nil {
		s.bucketService = bucketService.NewService(
			s.Config(),
			s.BucketRepository(),
		)
	}

	return s.bucketService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.Config(),
			s.WhiteListRepository(ctx),
			s.BlackListRepository(ctx),
			s.BucketService(),
		)
	}

	return s.accessService
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(
			s.Config(),
			s.AccessService(ctx),
			s.BucketService(),
		)
	}

	return s.accessImpl
}
