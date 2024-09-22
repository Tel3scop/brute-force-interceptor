package bucket

import (
	"context"
	"testing"

	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	serviceMocks "github.com/Tel3scop/brute-force-interceptor/tests/service"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestResetBuckets(t *testing.T) {
	tests := []struct {
		name         string
		bucketKeys   []string
		deleteErrors []error
		expectedErr  error
	}{
		{
			name:         "All deletions succeed",
			bucketKeys:   []string{"bucket1", "bucket2"},
			deleteErrors: []error{nil, nil},
			expectedErr:  nil,
		},
		{
			name:         "One deletion fails",
			bucketKeys:   []string{"bucket1", "bucket2"},
			deleteErrors: []error{nil, assert.AnError},
			expectedErr:  status.Errorf(codes.Internal, "failed to remove bucket2: %v", assert.AnError),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repositoryMocks.NewMockBucketRepository(ctrl)

			mockRepo.EXPECT().
				UsePipeline(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
					mockPipe := serviceMocks.NewMockPipeliner(ctrl)
					return fn(mockPipe)
				})

			for i, key := range tt.bucketKeys {
				mockRepo.EXPECT().Delete(gomock.Any(), key).Return(tt.deleteErrors[i])
			}

			s := &serv{bucketRepository: mockRepo}
			err := s.ResetBuckets(context.Background(), tt.bucketKeys)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
