package access

import (
	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/service"
	"github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// Implementation структура для работы с хэндерами
type Implementation struct {
	access_v1.UnimplementedAntiBruteforceServer
	accessService service.AccessService
	bucketService service.BucketService
	cfg           *config.Config
}

// NewImplementation новый экземпляр структуры Implementation
func NewImplementation(cfg *config.Config, accessService service.AccessService, bucketService service.BucketService) *Implementation {
	return &Implementation{
		cfg:           cfg,
		accessService: accessService,
		bucketService: bucketService,
	}
}
