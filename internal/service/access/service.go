package access

import (
	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	"github.com/Tel3scop/brute-force-interceptor/internal/service"
)

type serv struct {
	whiteListRepository repository.WhiteListRepository
	blackListRepository repository.BlackListRepository
	bucketService       service.BucketService
	cfg                 *config.Config
}

// NewService функция возвращает новый сервис управления доступом
func NewService(
	cfg *config.Config,
	whiteListRepository repository.WhiteListRepository,
	blackListRepository repository.BlackListRepository,
	bucketService service.BucketService,
) service.AccessService {
	return &serv{
		cfg:                 cfg,
		whiteListRepository: whiteListRepository,
		blackListRepository: blackListRepository,
		bucketService:       bucketService,
	}
}
