package bucket

import (
	"context"
	"errors"
	"testing"
	"time"

	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetRequestTimestamps(t *testing.T) {
	tests := []struct {
		name        string
		bucketKey   string
		timestamps  []time.Time
		expectedErr error
	}{
		{
			name:        "Successful retrieval of timestamps",
			bucketKey:   "bucket1",
			timestamps:  []time.Time{time.Now(), time.Now().Add(-time.Minute)},
			expectedErr: nil,
		},
		{
			name:        "Error retrieving timestamps",
			bucketKey:   "bucket2",
			timestamps:  nil,
			expectedErr: errors.New("retrieval error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repositoryMocks.NewMockBucketRepository(ctrl)

			mockRepo.EXPECT().GetRequestTimestamps(gomock.Any(), tt.bucketKey).Return(tt.timestamps, tt.expectedErr)

			s := &serv{bucketRepository: mockRepo}
			result, err := s.GetRequestTimestamps(context.Background(), tt.bucketKey)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.timestamps, result)
		})
	}
}
