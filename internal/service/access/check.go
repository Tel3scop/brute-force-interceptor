package access

import (
	"context"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	"github.com/Tel3scop/helpers/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Check проверяет возможность доступа.
func (s *serv) Check(ctx context.Context, auth model.Auth) error {
	ok, err := s.whiteListRepository.IsInList(ctx, auth.IP)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check whitelist: %v", err)
	}
	if ok {
		return nil
	}

	ok, err = s.blackListRepository.IsInList(ctx, auth.IP)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check blacklist: %v", err)
	}

	if ok {
		logger.Error("permission denied")
		return status.Errorf(codes.PermissionDenied, "permission denied")
	}

	now := time.Now()
	if err := s.bucketService.AddRequestTimestamps(ctx, auth, now); err != nil {
		return err
	}

	err = s.checkLimits(ctx, auth, now)
	if err != nil {
		return err
	}

	return nil
}
