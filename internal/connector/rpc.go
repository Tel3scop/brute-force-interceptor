package connector

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGRPCConnection(address string) *grpc.ClientConn {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

// AntiBruteforceClient is a client for BruteForceInterceptor.
type AntiBruteforceClient interface {
	ResetBucket(ctx context.Context, login, ip string) error
	AddToBlacklist(ctx context.Context, subnet string) error
	RemoveFromBlacklist(ctx context.Context, subnet string) error
	AddToWhitelist(ctx context.Context, subnet string) error
	RemoveFromWhitelist(ctx context.Context, subnet string) error
}
