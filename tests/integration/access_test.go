//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AccessSuite struct {
	suite.Suite
	ctx          context.Context
	accessClient access_v1.AntiBruteforceClient
	subnet       string
}

func (s *AccessSuite) SetupSuite() {
	const Localhost = "localhost:50051"
	host := os.Getenv("GRPC_ADDRESS")
	if host == "" {
		host = Localhost
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	s.ctx = context.Background()
	s.accessClient = access_v1.NewAntiBruteforceClient(conn)
}

func (s *AccessSuite) SetupTest() {
	seed := time.Now().UnixNano()
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
