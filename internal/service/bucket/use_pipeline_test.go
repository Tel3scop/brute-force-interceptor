package bucket

import (
	"context"
	"errors"
	"testing"

	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUsePipeline_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBucketRepo := repositoryMocks.NewMockBucketRepository(ctrl)
	s := &serv{bucketRepository: mockBucketRepo}

	mockBucketRepo.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(pipe redis.Pipeliner) error) error {
			client := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			})
			return fn(client.Pipeline())
		})

	err := s.UsePipeline(context.Background(), func(_ redis.Pipeliner) error {
		return nil
	})

	assert.NoError(t, err)
}

func TestUsePipeline_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBucketRepo := repositoryMocks.NewMockBucketRepository(ctrl)
	s := &serv{bucketRepository: mockBucketRepo}

	expectedError := errors.New("pipeline error")
	mockBucketRepo.EXPECT().UsePipeline(gomock.Any(), gomock.Any()).Return(expectedError)

	err := s.UsePipeline(context.Background(), func(_ redis.Pipeliner) error {
		return nil
	})

	assert.Equal(t, expectedError, err)
}
