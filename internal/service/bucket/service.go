package bucket

import (
	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	"github.com/Tel3scop/brute-force-interceptor/internal/service"
)

type serv struct {
	bucketRepository repository.BucketRepository
	cfg              *config.Config
}

// NewService функция возвращает новый сервис управления бакетами
func NewService(
	cfg *config.Config,
	bucketRepository repository.BucketRepository,
) service.BucketService {
	return &serv{
		cfg:              cfg,
		bucketRepository: bucketRepository,
	}
}
