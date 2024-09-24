package bucket

import (
	"context"
	"time"

	"github.com/Tel3scop/helpers/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serv) CheckBucketLimit(ctx context.Context, bucketKey string, time time.Time, limit int) error {
	timestamps, err := s.GetRequestTimestamps(ctx, bucketKey)
	if err != nil {
		return err
	}

	windowStart := time.Add(-s.cfg.Bucket.WindowSize)
	var count int
	for _, ts := range timestamps {
		if ts.After(windowStart) {
			count++
		}
	}

	if count > limit {
		logger.Error("Bucket limit exceeded", zap.Int("Count", count), zap.Int("Limit", limit))
		return status.Errorf(codes.OutOfRange, "request timestamps exceeds limit")
	}
	return nil
}
