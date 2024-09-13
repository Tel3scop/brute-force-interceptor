package access

import (
	"context"

	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// AddToWhitelist handler для добавления в белый список
func (i *Implementation) AddToWhitelist(ctx context.Context, request *accessAPI.AddToWhitelistRequest) (*accessAPI.AddToWhitelistResponse, error) {
	_, err := i.accessService.AddToWhiteList(ctx, request.GetSubnet())
	if err != nil {
		return nil, err
	}

	return &accessAPI.AddToWhitelistResponse{}, nil
}
