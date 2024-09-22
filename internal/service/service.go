package service

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/go-redis/redis/v8"
)

// BucketService интерфейс для управления бакетами.
type BucketService interface {
	Reset(ctx context.Context, bucket string) error
	ResetBuckets(ctx context.Context, bucketKeys []string) error
	AddRequestTimestamps(ctx context.Context, auth model.Auth, time time.Time) error
	GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error)
	UsePipeline(ctx context.Context, fn func(pipe redis.Pipeliner) error) error
	CheckBucketLimit(ctx context.Context, bucketKey string, time time.Time, limit int) error
}

// AccessService интерфейс для проверки доступа.
type AccessService interface {
	Check(ctx context.Context, auth model.Auth) error
	AddToBlackList(ctx context.Context, subnet string) (int64, error)
	AddToWhiteList(ctx context.Context, subnet string) (int64, error)
}
