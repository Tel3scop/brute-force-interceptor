package bucket

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (s *serv) UsePipeline(ctx context.Context, fn func(pipe redis.Pipeliner) error) error {
	return s.bucketRepository.UsePipeline(ctx, fn)
}
