package access

import (
	"context"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddToBlackList добавляет подсеть в черный список.
func (s *serv) AddToBlackList(ctx context.Context, subnet string) (int64, error) {
	blackListID, err := s.blackListRepository.Create(ctx, model.BlackList{Subnet: subnet})
	if err != nil {
		return 0, status.Errorf(codes.Internal, "can not create blacklist %s: %v", subnet, err)
	}

	return blackListID, nil
}
