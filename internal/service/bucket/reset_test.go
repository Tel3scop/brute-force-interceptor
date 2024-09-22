package bucket

import (
	"context"
	"errors"
	"testing"

	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestReset(t *testing.T) {
	tests := []struct {
		name        string
		bucket      string
		deleteError error
		expectedErr error
	}{
		{
			name:        "Successful deletion of bucket",
			bucket:      "bucket1",
			deleteError: nil,
			expectedErr: nil,
		},
		{
			name:        "Error deleting bucket",
			bucket:      "bucket2",
			deleteError: errors.New("deletion error"),
			expectedErr: status.Errorf(codes.Internal, "failed to delete bucket: %v", errors.New("deletion error")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repositoryMocks.NewMockBucketRepository(ctrl)

			mockRepo.EXPECT().Delete(gomock.Any(), tt.bucket).Return(tt.deleteError)

			s := &serv{bucketRepository: mockRepo}
			err := s.Reset(context.Background(), tt.bucket)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
