package access

import (
	"context"

	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// RemoveFromWhitelist handler для удаления из белого списка.
func (i *Implementation) RemoveFromWhitelist(
	ctx context.Context,
	request *accessAPI.RemoveFromWhitelistRequest,
) (*accessAPI.RemoveFromWhitelistResponse, error) {
	err := i.accessService.RemoveFromWhiteList(ctx, request.GetSubnet())
	if err != nil {
		return nil, err
	}

	return &accessAPI.RemoveFromWhitelistResponse{}, nil
}
