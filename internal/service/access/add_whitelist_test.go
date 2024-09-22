package access

import (
	"context"
	"errors"
	"testing"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddToWhiteList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositoryMocks.NewMockWhiteListRepository(ctrl)
	s := &serv{whiteListRepository: mockRepo}

	tests := []struct {
		name          string
		subnet        string
		mockReturnID  int64
		mockReturnErr error
		expectedID    int64
		expectedErr   error
	}{
		{
			name:          "Successful addition",
			subnet:        "192.168.1.0/24",
			mockReturnID:  1,
			mockReturnErr: nil,
			expectedID:    1,
			expectedErr:   nil,
		},
		{
			name:          "Repository error",
			subnet:        "192.168.1.0/24",
			mockReturnID:  0,
			mockReturnErr: errors.New("repository error"),
			expectedID:    0,
			expectedErr: status.Errorf(
				codes.Internal,
				"can not create whitelist %s: %v", "192.168.1.0/24",
				errors.New("repository error"),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().Create(gomock.Any(), model.WhiteList{Subnet: tt.subnet}).Return(tt.mockReturnID, tt.mockReturnErr)

			id, err := s.AddToWhiteList(context.Background(), tt.subnet)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
