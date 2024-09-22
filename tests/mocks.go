package mocks

//go:generate mockgen -source=../internal/service/service.go -destination=service/service.go -package=serviceMocks
//go:generate mockgen -source=../internal/repository/repository.go -destination=repository/repository.go -package=repositoryMocks
//go:generate mockgen -source=$GOPATH/pkg/mod/github.com/go-redis/redis/v8@v8.11.5/pipeline.go -destination=service/redis_mock.go -package=serviceMocks Pipeliner
