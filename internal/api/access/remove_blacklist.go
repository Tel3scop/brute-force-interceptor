package access

import (
	"context"

	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// RemoveFromBlacklist handler для удаления из черного списка.
func (i *Implementation) RemoveFromBlacklist(
	ctx context.Context,
	request *accessAPI.RemoveFromBlacklistRequest,
) (*accessAPI.RemoveFromBlacklistResponse, error) {
	_, err := i.accessService.AddToBlackList(ctx, request.GetSubnet())
	if err != nil {
		return nil, err
	}

	return &accessAPI.RemoveFromBlacklistResponse{}, nil
}
