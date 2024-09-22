package access

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	repositoryMocks "github.com/Tel3scop/brute-force-interceptor/tests/repository"
	serviceMocks "github.com/Tel3scop/brute-force-interceptor/tests/service"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setup(t *testing.T) (
	*serv,
	*repositoryMocks.MockWhiteListRepository,
	*repositoryMocks.MockBlackListRepository,
	*serviceMocks.MockBucketService,
) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockWhiteListRepo := repositoryMocks.NewMockWhiteListRepository(ctrl)
	mockBlackListRepo := repositoryMocks.NewMockBlackListRepository(ctrl)
	mockBucketService := serviceMocks.NewMockBucketService(ctrl)

	s := &serv{
		whiteListRepository: mockWhiteListRepo,
		blackListRepository: mockBlackListRepo,
		bucketService:       mockBucketService,
		cfg: &config.Config{
			Bucket: config.Bucket{
				LoginLimit:    5,
				PasswordLimit: 5,
				IPLimit:       5,
				WindowSize:    time.Minute,
			},
		},
	}

	return s, mockWhiteListRepo, mockBlackListRepo, mockBucketService
}

func TestCheck_IPInWhitelist(t *testing.T) {
	s, mockWhiteListRepo, _, _ := setup(t)

	auth := model.Auth{IP: "192.168.1.1"}
	mockWhiteListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(true, nil)

	err := s.Check(context.Background(), auth)

	assert.NoError(t, err)
}

func TestCheck_IPInBlacklist(t *testing.T) {
	s, mockWhiteListRepo, mockBlackListRepo, _ := setup(t)

	auth := model.Auth{IP: "192.168.1.1"}
	mockWhiteListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBlackListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(true, nil)

	err := s.Check(context.Background(), auth)

	assert.Equal(t, status.Errorf(codes.PermissionDenied, "permission denied"), err)
}

func TestCheck_WhitelistCheckError(t *testing.T) {
	s, mockWhiteListRepo, _, _ := setup(t)

	auth := model.Auth{IP: "192.168.1.1"}
	mockWhiteListRepo.EXPECT().
		IsInList(gomock.Any(), auth.IP).
		Return(false, errors.New("whitelist check error"))

	err := s.Check(context.Background(), auth)

	assert.Equal(
		t,
		status.Errorf(codes.Internal, "failed to check whitelist: %v", errors.New("whitelist check error")),
		err,
	)
}

func TestCheck_BlacklistCheckError(t *testing.T) {
	s, mockWhiteListRepo, mockBlackListRepo, _ := setup(t)

	auth := model.Auth{IP: "192.168.1.1"}
	mockWhiteListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBlackListRepo.EXPECT().
		IsInList(gomock.Any(), auth.IP).
		Return(false, errors.New("blacklist check error"))

	err := s.Check(context.Background(), auth)

	assert.Equal(
		t,
		status.Errorf(codes.Internal, "failed to check blacklist: %v", errors.New("blacklist check error")),
		err,
	)
}

func TestCheck_AddRequestTimestampsError(t *testing.T) {
	s, mockWhiteListRepo, mockBlackListRepo, mockBucketService := setup(t)

	auth := model.Auth{IP: "192.168.1.1"}
	mockWhiteListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBlackListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBucketService.EXPECT().
		AddRequestTimestamps(gomock.Any(), auth, gomock.Any()).
		Return(errors.New("add request error"))

	err := s.Check(context.Background(), auth)

	assert.Equal(t, errors.New("add request error"), err)
}

func TestCheck_SuccessfulCheck(t *testing.T) {
	s, mockWhiteListRepo, mockBlackListRepo, mockBucketService := setup(t)

	auth := model.Auth{
		Login:    "user1",
		Password: "pass1",
		IP:       "192.168.1.1",
	}
	mockWhiteListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBlackListRepo.EXPECT().IsInList(gomock.Any(), auth.IP).Return(false, nil)
	mockBucketService.EXPECT().AddRequestTimestamps(gomock.Any(), auth, gomock.Any()).Return(nil)
	mockBucketService.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, f func(redis.Pipeliner) error) error {
			mockPipeliner := serviceMocks.NewMockPipeliner(gomock.NewController(t))
			return f(mockPipeliner)
		})

	mockBucketService.EXPECT().CheckBucketLimit(gomock.Any(), "login:user1", gomock.Any(), 5).AnyTimes().Return(nil)
	mockBucketService.EXPECT().CheckBucketLimit(gomock.Any(), "password:pass1", gomock.Any(), 5).AnyTimes().Return(nil)
	mockBucketService.EXPECT().CheckBucketLimit(gomock.Any(), "ip:192.168.1.1", gomock.Any(), 5).AnyTimes().Return(nil)

	err := s.Check(context.Background(), auth)

	assert.NoError(t, err)
}
