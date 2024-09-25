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

type WhiteListSuite struct {
	suite.Suite
	ctx          context.Context
	clientConn   *grpc.ClientConn
	accessClient access_v1.AntiBruteforceClient
	subnet       string
}

func (s *WhiteListSuite) SetupSuite() {
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		host = "127.0.0.1:50051"
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	s.ctx = context.Background()
	s.accessClient = access_v1.NewAntiBruteforceClient(conn)
}

func (s *WhiteListSuite) SetupTest() {
	var seed = time.Now().UnixNano()
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
