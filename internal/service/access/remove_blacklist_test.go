package access

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

func TestRemoveBlackList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositoryMocks.NewMockBlackListRepository(ctrl)
	s := &serv{blackListRepository: mockRepo}

	tests := []struct {
		name          string
		subnet        string
		mockReturnErr error
		expectedErr   error
	}{
		{
			name:          "Successful remove",
			subnet:        "192.168.1.0/24",
			mockReturnErr: nil,
			expectedErr:   nil,
		},
		{
			name:          "Repository error",
			subnet:        "192.168.1.0/24",
			mockReturnErr: errors.New("repository error"),
			expectedErr: status.Errorf(
				codes.Internal,
				"can not remove blacklist %s: %v", "192.168.1.0/24",
				errors.New("repository error"),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().Delete(gomock.Any(), tt.subnet).Return(tt.mockReturnErr)

			err := s.RemoveFromBlackList(context.Background(), tt.subnet)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
