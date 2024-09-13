package access

import (
	"context"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddToWhiteList добавляет подсеть в белый список
func (s *serv) AddToWhiteList(ctx context.Context, subnet string) (int64, error) {
	whiteListID, err := s.whiteListRepository.Create(ctx, model.WhiteList{Subnet: subnet})
	if err != nil {
		return 0, status.Errorf(codes.Internal, "can not create whitelist %s: %v", subnet, err)
	}

	return whiteListID, nil
}
