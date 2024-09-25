package integration

import (
	"context"
	"fmt"
	"github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"
)

type AccessSuite struct {
	suite.Suite
	ctx          context.Context
	clientConn   *grpc.ClientConn
	accessClient access_v1.AntiBruteforceClient
	subnet       string
}

func (s *AccessSuite) SetupSuite() {
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		host = "127.0.0.1:50051"
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	s.ctx = context.Background()
	s.accessClient = access_v1.NewAntiBruteforceClient(conn)
}

func (s *AccessSuite) SetupTest() {
	var seed = time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	ip := gofakeit.IPv4Address()
	mask := fmt.Sprintf("/%d", gofakeit.Number(0, 32))
	_, ipNet, err := net.ParseCIDR(ip + mask)
	s.Require().NoError(err)

	s.subnet = ipNet.String()

	s.T().Log("seed:", seed)

}

func TestAccessSuite(t *testing.T) {
	suite.Run(t, new(AccessSuite))
}

func (s *AccessSuite) TestTryAuthNegative() {
	req := &access_v1.AuthRequest{
		Login:    "invalid-login",
		Password: "invalid-password",
		Ip:       "invalid-ip",
	}
	resp, err := s.accessClient.TryAuth(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
	s.Require().Nil(resp, status.Code(err))
}

func (s *AccessSuite) TestResetBucket() {
	req := &access_v1.ResetBucketRequest{
		Login: gofakeit.Username(),
		Ip:    gofakeit.IPv4Address(),
	}
	_, err := s.accessClient.ResetBucket(s.ctx, req)
	s.Require().NoError(err)
}

func (s *AccessSuite) TestResetBucketNegative() {
	req := &access_v1.ResetBucketRequest{
		Login: "invalid-login",
		Ip:    "invalid-ip",
	}
	_, err := s.accessClient.ResetBucket(s.ctx, req)
	s.Require().NoError(err)
}
