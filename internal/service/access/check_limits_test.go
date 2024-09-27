package access

import (
	"context"
	"testing"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/config"
	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	serviceMocks "github.com/Tel3scop/brute-force-interceptor/tests/service"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCheckLimits_AllBucketsWithinLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBucketService := serviceMocks.NewMockBucketService(ctrl)

	serv := &serv{
		bucketService: mockBucketService,
		cfg: &config.Config{
			Bucket: config.Bucket{
				LoginLimit:    5,
				PasswordLimit: 5,
				IPLimit:       5,
			},
		},
	}

	ctx := context.Background()
	auth := model.Auth{
		Login:    "testLogin",
		Password: "testPassword",
		IP:       "192.168.1.1",
	}
	timeNow := time.Now()

	mockBucketService.EXPECT().
		UsePipeline(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(ctrl)
			return fn(mockPipe)
		})

	mockBucketService.EXPECT().
		CheckBucketLimit(ctx, auth.LoginBucket(), timeNow, serv.cfg.Bucket.LoginLimit).
		AnyTimes().
		Return(nil)
	mockBucketService.EXPECT().
		CheckBucketLimit(ctx, auth.PasswordBucket(), timeNow, serv.cfg.Bucket.PasswordLimit).
		AnyTimes().
		Return(nil)
	mockBucketService.EXPECT().
		CheckBucketLimit(ctx, auth.IPBucket(), timeNow, serv.cfg.Bucket.IPLimit).
		AnyTimes().
		Return(nil)

	err := serv.checkLimits(ctx, auth, timeNow)

	assert.NoError(t, err)
}

func TestCheckLimits_LoginBucketExceedsLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBucketService := serviceMocks.NewMockBucketService(ctrl)
	cfg := &config.Config{
		Bucket: config.Bucket{
			LoginLimit:    5,
			PasswordLimit: 5,
			IPLimit:       5,
		},
	}
	serv := &serv{
		bucketService: mockBucketService,
		cfg:           cfg,
	}

	ctx := context.Background()
	auth := model.Auth{
		Login:    "testLogin",
		Password: "testPassword",
		IP:       "192.168.1.1",
	}
	timeNow := time.Now()

	mockBucketService.EXPECT().
		UsePipeline(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, fn func(redis.Pipeliner) error) error {
			mockPipe := serviceMocks.NewMockPipeliner(ctrl)

			return fn(mockPipe)
		})

	mockBucketService.EXPECT().CheckBucketLimit(ctx, auth.LoginBucket(), timeNow, 5).AnyTimes().Return(assert.AnError)
	mockBucketService.EXPECT().CheckBucketLimit(ctx, auth.PasswordBucket(), timeNow, 5).AnyTimes().Return(nil)
	mockBucketService.EXPECT().CheckBucketLimit(ctx, auth.IPBucket(), timeNow, 5).AnyTimes().Return(nil)

	err := serv.checkLimits(ctx, auth, timeNow)

	assert.Equal(t, assert.AnError, err)
}
