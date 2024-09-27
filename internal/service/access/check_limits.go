package access

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/go-redis/redis/v8"
)

func (s *serv) checkLimits(ctx context.Context, auth model.Auth, time time.Time) error {
	return s.bucketService.UsePipeline(ctx, func(_ redis.Pipeliner) error {
		limits := map[string]int{
			auth.LoginBucket():    s.cfg.Bucket.LoginLimit,
			auth.PasswordBucket(): s.cfg.Bucket.PasswordLimit,
			auth.IPBucket():       s.cfg.Bucket.IPLimit,
		}

		for bucketKey, limit := range limits {
			err := s.bucketService.CheckBucketLimit(ctx, bucketKey, time, limit)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
