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

type BlackListSuite struct {
	suite.Suite
	ctx          context.Context
	accessClient access_v1.AntiBruteforceClient
	subnet       string
}

func (s *BlackListSuite) SetupSuite() {
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

func (s *BlackListSuite) SetupTest() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	ip := gofakeit.IPv4Address()
	mask := fmt.Sprintf("/%d", gofakeit.Number(0, 32))
	_, ipNet, err := net.ParseCIDR(ip + mask)
	s.Require().NoError(err)

	s.subnet = ipNet.String()

	s.T().Log("seed:", seed)
}

func TestBlacklistSuite(t *testing.T) {
	suite.Run(t, new(BlackListSuite))
}

func (s *BlackListSuite) TestAddToBlacklist() {
	req := &access_v1.AddToBlacklistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.AddToBlacklist(s.ctx, req)
	s.Require().NoError(err)
}

func (s *BlackListSuite) TestAddToBlacklistNegative() {
	req := &access_v1.AddToBlacklistRequest{
		Subnet: "invalid-subnet",
	}
	_, err := s.accessClient.AddToBlacklist(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
}

func (s *BlackListSuite) TestRemoveFromBlacklist() {
	addReq := &access_v1.AddToBlacklistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.AddToBlacklist(s.ctx, addReq)
	s.Require().NoError(err)

	removeReq := &access_v1.RemoveFromBlacklistRequest{
		Subnet: s.subnet,
	}
	_, err = s.accessClient.RemoveFromBlacklist(s.ctx, removeReq)
	s.Require().NoError(err)
}

func (s *BlackListSuite) TestRemoveFromBlacklistNegative() {
	removeReq := &access_v1.RemoveFromBlacklistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.RemoveFromBlacklist(s.ctx, removeReq)
	s.Require().Nil(err)
}
