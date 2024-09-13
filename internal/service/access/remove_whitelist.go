package access

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RemoveFromWhiteList удаляет подсеть из белого списка
func (s *serv) RemoveFromWhiteList(ctx context.Context, subnet string) error {
	err := s.whiteListRepository.Delete(ctx, subnet)
	if err != nil {
		return status.Errorf(codes.Internal, "can not remove whitelist %s: %v", subnet, err)
	}

	return nil
}
