package bucket

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddRequestTimestamps метод устанавливает последнее время использования бакета.
func (s *serv) AddRequestTimestamps(ctx context.Context, auth model.Auth, time time.Time) error {
	return s.bucketRepository.UsePipeline(ctx, func(_ redis.Pipeliner) error {
		if err := s.bucketRepository.AddRequestTimestamp(ctx, auth.LoginBucket(), time, s.cfg.Bucket.TTL); err != nil {
			return status.Errorf(codes.Internal, "failed to add login request timestamp: %v", err)
		}

		if err := s.bucketRepository.AddRequestTimestamp(ctx, auth.PasswordBucket(), time, s.cfg.Bucket.TTL); err != nil {
			return status.Errorf(codes.Internal, "failed to add password request timestamp: %v", err)
		}

		if err := s.bucketRepository.AddRequestTimestamp(ctx, auth.IPBucket(), time, s.cfg.Bucket.TTL); err != nil {
			return status.Errorf(codes.Internal, "failed to add ip request timestamp: %v", err)
		}

		return nil
	})
}
