package repository

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/go-redis/redis/v8"
)

// WhiteListRepository интерфейс репозитория белых списков
type WhiteListRepository interface {
	Create(ctx context.Context, dto model.WhiteList) (int64, error)
	Delete(ctx context.Context, subnet string) error
	IsInList(ctx context.Context, ip string) (bool, error)
}

// BlackListRepository интерфейс репозитория черных списков
type BlackListRepository interface {
	Create(ctx context.Context, dto model.BlackList) (int64, error)
	Delete(ctx context.Context, subnet string) error
	IsInList(ctx context.Context, ip string) (bool, error)
}

// BucketRepository интерфейс репозитория для управления бакетами
type BucketRepository interface {
	AddRequestTimestamp(ctx context.Context, bucketKey string, timestamp time.Time, ttl time.Duration) error
	GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error)
	Delete(ctx context.Context, bucketKey string) error
	UsePipeline(ctx context.Context, fn func(pipe redis.Pipeliner) error) error
}
