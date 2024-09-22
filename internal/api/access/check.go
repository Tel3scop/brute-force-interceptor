package access

import (
	"context"

	"github.com/Tel3scop/brute-force-interceptor/internal/model"
	accessAPI "github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
)

// TryAuth handler для проверки авторизации.
func (i *Implementation) TryAuth(ctx context.Context, request *accessAPI.AuthRequest) (*accessAPI.AuthResponse, error) {
	auth := model.Auth{
		Login:    request.GetLogin(),
		Password: request.GetPassword(),
		IP:       request.GetIp(),
	}

	err := i.accessService.Check(ctx, auth)
	if err != nil {
		return &accessAPI.AuthResponse{Ok: false}, err
	}

	return &accessAPI.AuthResponse{Ok: true}, nil
}
