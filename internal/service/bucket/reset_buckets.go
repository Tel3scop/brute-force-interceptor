package bucket

import (
	"context"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ResetBuckets метод устанавливает последнее время использования бакета.
func (s *serv) ResetBuckets(ctx context.Context, bucketKeys []string) error {
	return s.bucketRepository.UsePipeline(ctx, func(_ redis.Pipeliner) error {
		for _, bucketKey := range bucketKeys {
			if err := s.bucketRepository.Delete(ctx, bucketKey); err != nil {
				return status.Errorf(codes.Internal, "failed to remove %s: %v", bucketKey, err)
			}
		}
		return nil
	})
}
