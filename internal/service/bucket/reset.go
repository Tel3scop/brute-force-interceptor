package bucket

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Reset метод сбрасывает счетчик бакета
func (s *serv) Reset(ctx context.Context, bucket string) error {
	err := s.bucketRepository.Delete(ctx, bucket)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to delete bucket: %v", err)
	}

	return nil
}
