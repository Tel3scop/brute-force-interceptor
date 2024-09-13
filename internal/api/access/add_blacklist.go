package access

import (
	"context"

	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// AddToBlacklist handler для добавления в черный список
func (i *Implementation) AddToBlacklist(ctx context.Context, request *accessAPI.AddToBlacklistRequest) (*accessAPI.AddToBlacklistResponse, error) {
	_, err := i.accessService.AddToBlackList(ctx, request.GetSubnet())
	if err != nil {
		return nil, err
	}

	return &accessAPI.AddToBlacklistResponse{}, nil
}
