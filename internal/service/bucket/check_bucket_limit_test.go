package bucket

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCheckBucketLimit(t *testing.T) {
	tests := []struct {
		name        string
		bucketKey   string
		time        time.Time
		limit       int
		timestamps  []time.Time
		getError    error
		expectedErr error
	}{
		{
			name:        "Successful check within limit",
			bucketKey:   "bucket1",
			time:        time.Now(),
			limit:       5,
			timestamps:  []time.Time{time.Now(), time.Now().Add(-time.Minute)},
			getError:    nil,
			expectedErr: nil,
		},
		{
			name:        "Error getting request timestamps",
			bucketKey:   "bucket2",
			time:        time.Now(),
			limit:       5,
			timestamps:  nil,
			getError:    errors.New("get error"),
			expectedErr: errors.New("get error"),
		},
		{
			name:        "Request timestamps exceeds limit",
			bucketKey:   "bucket3",
			time:        time.Now(),
			limit:       1,
			timestamps:  []time.Time{time.Now(), time.Now().Add(-time.Minute)},
			getError:    errors.New("get error"),
			expectedErr: errors.New("get error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repositoryMocks.NewMockBucketRepository(ctrl)

			mockRepo.EXPECT().GetRequestTimestamps(gomock.Any(), tt.bucketKey).Return(tt.timestamps, tt.getError)

			s := &serv{
				bucketRepository: mockRepo,
				cfg: &config.Config{
					Bucket: config.Bucket{
						WindowSize: time.Minute,
					},
				},
			}

			err := s.CheckBucketLimit(context.Background(), tt.bucketKey, tt.time, tt.limit)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
