package bucket

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/repository"
	"github.com/go-redis/redis/v8"
)

type repo struct {
	client *redis.Client
}

// NewRepository создание репозитория
func NewRepository(client *redis.Client) repository.BucketRepository {
	return &repo{client: client}
}

//
//func (r *repo) IncrementRequestCount(ctx context.Context, bucketKey string) (int64, error) {
//	result, err := r.client.HIncrBy(ctx, bucketKey, "request_count", 1).Result()
//
//	if err != nil {
//		logger.Error("failed to increment count", zap.String("key", bucketKey), zap.Error(err))
//		return 0, err
//	}
//
//	return result, nil
//}
//
//func (r *repo) GetLastRequestTime(ctx context.Context, bucketKey string) (string, error) {
//	result, err := r.client.HGet(ctx, bucketKey, "last_request_time").Result()
//	if err != nil {
//		logger.Error("failed to get last request time", zap.String("key", bucketKey), zap.Error(err))
//		return "", err
//	}
//
//	return result, nil
//}
//
//func (r *repo) SetLastRequestTime(ctx context.Context, bucketKey, lastRequestTime string) error {
//	err := r.client.HSet(ctx, bucketKey, "last_request_time", lastRequestTime).Err()
//	if err != nil {
//		logger.Error("failed to set last request time", zap.String("key", bucketKey), zap.Error(err))
//		return err
//	}
//
//	return nil
//}
//
//func (r *repo) Delete(ctx context.Context, bucketKey string) error {
//	err := r.client.Del(ctx, bucketKey).Err()
//	if err != nil {
//		logger.Error("failed to delete bucket", zap.String("key", bucketKey), zap.Error(err))
//		return err
//	}
//
//	return nil
//}
//
//func (r *repo) ScanKeys(ctx context.Context, pattern string) ([]string, error) {
//	var keys []string
//	var cursor uint64
//
//	for {
//		var err error
//		var result []string
//		result, cursor, err = r.client.Scan(ctx, cursor, pattern, 10).Result()
//		if err != nil {
//			logger.Error("failed to scan keys", zap.String("pattern", pattern), zap.Error(err))
//			return nil, err
//		}
//
//		keys = append(keys, result...)
//
//		if cursor == 0 {
//			break
//		}
//	}
//
//	return keys, nil
//}
//
//func (r *repo) AddRequestTimestamp(ctx context.Context, bucketKey string, timestamp time.Time) error {
//	timestamps, err := r.GetRequestTimestamps(ctx, bucketKey)
//	if err != nil {
//		return err
//	}
//
//	timestamps = append(timestamps, timestamp)
//	timestampsJSON, err := json.Marshal(timestamps)
//	if err != nil {
//		return err
//	}
//
//	return r.client.Set(ctx, bucketKey, timestampsJSON, 0).Err()
//}
//
//func (r *repo) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
//	timestampsJSON, err := r.client.Get(ctx, bucketKey).Result()
//	if err != nil && !errors.Is(err, redis.Nil) {
//		return nil, err
//	}
//
//	var timestamps []time.Time
//	if timestampsJSON != "" {
//		if err := json.Unmarshal([]byte(timestampsJSON), &timestamps); err != nil {
//			return nil, err
//		}
//	}
//
//	return timestamps, nil
//}

func (r *repo) AddRequestTimestamp(ctx context.Context, bucketKey string, timestamp time.Time, ttl time.Duration) error {
	timestamps, err := r.GetRequestTimestamps(ctx, bucketKey)
	if err != nil {
		return err
	}

	timestamps = append(timestamps, timestamp)
	timestampsJSON, err := json.Marshal(timestamps)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, bucketKey, timestampsJSON, ttl).Err()
}

func (r *repo) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
	timestampsJSON, err := r.client.Get(ctx, bucketKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	var timestamps []time.Time
	if timestampsJSON != "" {
		if err := json.Unmarshal([]byte(timestampsJSON), &timestamps); err != nil {
			return nil, err
		}
	}

	return timestamps, nil
}

func (r *repo) Delete(ctx context.Context, bucketKey string) error {
	return r.client.Del(ctx, bucketKey).Err()
}

func (r *repo) UsePipeline(ctx context.Context, fn func(pipe redis.Pipeliner) error) error {
	pipe := r.client.Pipeline()
	err := fn(pipe)
	if err != nil {
		return err
	}
	_, err = pipe.Exec(ctx)
	return err
}
