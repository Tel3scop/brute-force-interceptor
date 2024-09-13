package access

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RemoveFromBlackList удаляет подсеть из черного списка
func (s *serv) RemoveFromBlackList(ctx context.Context, subnet string) error {
	err := s.blackListRepository.Delete(ctx, subnet)
	if err != nil {
		return status.Errorf(codes.Internal, "can not remove blacklist %s: %v", subnet, err)
	}

	return nil
}
