package access

import (
	"context"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// ResetBucket handler для сброса бакета.
func (i *Implementation) ResetBucket(
	ctx context.Context,
	request *accessAPI.ResetBucketRequest,
) (*accessAPI.ResetBucketResponse, error) {
	auth := model.Auth{
		Login: request.Login,
		IP:    request.Ip,
	}

	buckets := append([]string{}, auth.LoginBucket(), auth.IPBucket())
	err := i.bucketService.ResetBuckets(ctx, buckets)
	if err != nil {
		return nil, err
	}

	return &accessAPI.ResetBucketResponse{}, nil
}
