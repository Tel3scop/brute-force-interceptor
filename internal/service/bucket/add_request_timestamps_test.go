package bucket

import (
	"context"
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

func setupAddRequestTimestampsTest(t *testing.T) (*serv, *repositoryMocks.MockBucketRepository, model.Auth, time.Time) {
	t.Helper()

	ctrl := gomock.NewController(t)

	mockRepo := repositoryMocks.NewMockBucketRepository(ctrl)
	auth := model.Auth{
		Login:    "user1",
		Password: "pass1",
		IP:       "192.168.1.1",
	}
	timeNow := time.Now()

	s := &serv{
		bucketRepository: mockRepo,
		cfg: &config.Config{
			Bucket: config.Bucket{
				TTL: time.Minute,
			},
		},
	}

	return s, mockRepo, auth, timeNow
}

func TestAddRequestTimestamps_AllTimestampsAddedSuccessfully(t *testing.T) {
	s, mockRepo, auth, timeNow := setupAddRequestTimestampsTest(t)

	mockRepo.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(gomock.NewController(t))
			return fn(mockPipe)
		})

	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.LoginBucket(), timeNow, gomock.Any()).Return(nil)
	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.PasswordBucket(), timeNow, gomock.Any()).Return(nil)
	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.IPBucket(), timeNow, gomock.Any()).Return(nil)

	err := s.AddRequestTimestamps(context.Background(), auth, timeNow)

	assert.NoError(t, err)
}

func TestAddRequestTimestamps_ErrorAddingLoginTimestamp(t *testing.T) {
	s, mockRepo, auth, timeNow := setupAddRequestTimestampsTest(t)

	mockRepo.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(gomock.NewController(t))
			return fn(mockPipe)
		})

	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.LoginBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(assert.AnError)
	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.PasswordBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(nil)
	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.IPBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(nil)

	err := s.AddRequestTimestamps(context.Background(), auth, timeNow)

	expectedErr := status.Errorf(codes.Internal, "failed to add login request timestamp: %v", assert.AnError)
	assert.EqualError(t, err, expectedErr.Error())
}

func TestAddRequestTimestamps_ErrorAddingPasswordTimestamp(t *testing.T) {
	s, mockRepo, auth, timeNow := setupAddRequestTimestampsTest(t)

	mockRepo.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(gomock.NewController(t))
			return fn(mockPipe)
		})

	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.LoginBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(nil)
	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.PasswordBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(assert.AnError)
	mockRepo.EXPECT().
		AddRequestTimestamp(gomock.Any(), auth.IPBucket(), timeNow, gomock.Any()).
		AnyTimes().
		Return(nil)

	err := s.AddRequestTimestamps(context.Background(), auth, timeNow)

	expectedErr := status.Errorf(codes.Internal, "failed to add password request timestamp: %v", assert.AnError)
	assert.EqualError(t, err, expectedErr.Error())
}

func TestAddRequestTimestamps_ErrorAddingIpTimestamp(t *testing.T) {
	s, mockRepo, auth, timeNow := setupAddRequestTimestampsTest(t)

	mockRepo.EXPECT().
		UsePipeline(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(gomock.NewController(t))
			return fn(mockPipe)
		})

	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.LoginBucket(), timeNow, gomock.Any()).Return(nil)
	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.PasswordBucket(), timeNow, gomock.Any()).Return(nil)
	mockRepo.EXPECT().AddRequestTimestamp(gomock.Any(), auth.IPBucket(), timeNow, gomock.Any()).Return(assert.AnError)

	err := s.AddRequestTimestamps(context.Background(), auth, timeNow)

	expectedErr := status.Errorf(codes.Internal, "failed to add ip request timestamp: %v", assert.AnError)
	assert.EqualError(t, err, expectedErr.Error())
}
