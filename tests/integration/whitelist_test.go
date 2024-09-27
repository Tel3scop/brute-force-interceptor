//go:build integration

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

type WhiteListSuite struct {
	suite.Suite
	ctx          context.Context
	accessClient access_v1.AntiBruteforceClient
	subnet       string
}

func (s *WhiteListSuite) SetupSuite() {
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

func (s *WhiteListSuite) SetupTest() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	ip := gofakeit.IPv4Address()
	mask := fmt.Sprintf("/%d", gofakeit.Number(0, 32))
	_, ipNet, err := net.ParseCIDR(ip + mask)
	s.Require().NoError(err)

	s.subnet = ipNet.String()

	s.T().Log("seed:", seed)
}

func TestWhiteListSuite(t *testing.T) {
	suite.Run(t, new(WhiteListSuite))
}

func (s *WhiteListSuite) TestAddToWhitelist() {
	req := &access_v1.AddToWhitelistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.AddToWhitelist(s.ctx, req)
	s.Require().NoError(err)
}

func (s *WhiteListSuite) TestAddToWhitelistNegative() {
	req := &access_v1.AddToWhitelistRequest{
		Subnet: "invalid-subnet",
	}
	_, err := s.accessClient.AddToWhitelist(s.ctx, req)
	s.Require().Error(err)
	s.Require().Equal(codes.Internal, status.Code(err))
}

func (s *WhiteListSuite) TestRemoveFromWhitelist() {
	addReq := &access_v1.AddToWhitelistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.AddToWhitelist(s.ctx, addReq)
	s.Require().NoError(err)

	removeReq := &access_v1.RemoveFromWhitelistRequest{
		Subnet: s.subnet,
	}
	_, err = s.accessClient.RemoveFromWhitelist(s.ctx, removeReq)
	s.Require().NoError(err)
}

func (s *WhiteListSuite) TestRemoveFromWhitelistNegative() {
	removeReq := &access_v1.RemoveFromWhitelistRequest{
		Subnet: s.subnet,
	}
	_, err := s.accessClient.RemoveFromWhitelist(s.ctx, removeReq)
	s.Require().Nil(err)
}
