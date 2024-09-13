package access

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serv) checkLimits(ctx context.Context, auth model.Auth, time time.Time) error {
	return s.bucketService.UsePipeline(ctx, func(pipe redis.Pipeliner) error {
		limits := map[string]int{
			auth.LoginBucket():    s.cfg.Bucket.LoginLimit,
			auth.PasswordBucket(): s.cfg.Bucket.PasswordLimit,
			auth.IpBucket():       s.cfg.Bucket.IPLimit,
		}

		for bucketKey, limit := range limits {
			timestamps, err := s.bucketService.GetRequestTimestamps(ctx, bucketKey)
			if err != nil {
				return status.Errorf(codes.Internal, "failed to get request timestamps: %v", err)
			}

			windowStart := time.Add(-s.cfg.Bucket.WindowSize)
			var count int
			for _, ts := range timestamps {
				if ts.After(windowStart) {
					count++
				}
			}

			if count > limit {
				return status.Errorf(codes.OutOfRange, "request timestamps exceeds limit")
			}
		}

		return nil
	})
}
