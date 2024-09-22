package access

import (
	"context"

	"github.com/Tel3scop/brute-force-interceptor/internal/connector"
	"github.com/Tel3scop/brute-force-interceptor/pkg/access_v1"
	"google.golang.org/grpc"
)

// Client экземпляр.
type Client struct {
	client access_v1.AntiBruteforceClient
}

var _ connector.AntiBruteforceClient = (*Client)(nil)

// New создает новый экземпляр клиента.
func New(conn *grpc.ClientConn) (*Client, error) {
	return &Client{
		client: access_v1.NewAntiBruteforceClient(conn),
	}, nil
}

func (c *Client) ResetBucket(ctx context.Context, login, ip string) error {
	_, err := c.client.ResetBucket(ctx, &access_v1.ResetBucketRequest{Login: login, Ip: ip})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddToBlacklist(ctx context.Context, subnet string) error {
	_, err := c.client.AddToBlacklist(ctx, &access_v1.AddToBlacklistRequest{Subnet: subnet})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveFromBlacklist(ctx context.Context, subnet string) error {
	_, err := c.client.RemoveFromBlacklist(ctx, &access_v1.RemoveFromBlacklistRequest{Subnet: subnet})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddToWhitelist(ctx context.Context, subnet string) error {
	_, err := c.client.AddToWhitelist(ctx, &access_v1.AddToWhitelistRequest{Subnet: subnet})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveFromWhitelist(ctx context.Context, subnet string) error {
	_, err := c.client.RemoveFromWhitelist(ctx, &access_v1.RemoveFromWhitelistRequest{Subnet: subnet})
	if err != nil {
		return err
	}

	return nil
}
