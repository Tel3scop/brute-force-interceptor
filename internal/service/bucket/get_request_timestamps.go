package bucket

import (
	"context"
	"time"
)

func (s *serv) GetRequestTimestamps(ctx context.Context, bucketKey string) ([]time.Time, error) {
	return s.bucketRepository.GetRequestTimestamps(ctx, bucketKey)
}
