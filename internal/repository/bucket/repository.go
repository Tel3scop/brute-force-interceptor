package bucket

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	"github.com/Tel3scop/helpers/logger"
	"github.com/go-redis/redis/v8"
)

type repo struct {
	client *redis.Client
}

// NewRepository создание репозитория.
func NewRepository(client *redis.Client) repository.BucketRepository {
	return &repo{client: client}
}

func (r *repo) AddRequestTimestamp(ctx context.Context, bucketKey string, tStamp time.Time, ttl time.Duration) error {
	timestamps, err := r.GetRequestTimestamps(ctx, bucketKey)
	if err != nil {
		return err
	}

	timestamps = append(timestamps, tStamp)
	timestampsJSON, err := json.Marshal(timestamps)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return r.client.Set(ctx, bucketKey, timestampsJSON, ttl).Err()
}

func (r *repo) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
	timestampsJSON, err := r.client.Get(ctx, bucketKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Error(err.Error())
		return nil, err
	}

	var timestamps []time.Time
	if timestampsJSON != "" {
		if err := json.Unmarshal([]byte(timestampsJSON), &timestamps); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return timestamps, nil
}

func (r *repo) Delete(ctx context.Context, bucketKey string) error {
	err := r.client.Del(ctx, bucketKey).Err()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *repo) UsePipeline(ctx context.Context, fn func(pipe redis.Pipeliner) error) error {
	pipe := r.client.Pipeline()
	err := fn(pipe)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
