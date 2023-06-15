package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	client *rpc.Client
}

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	c, err := rpc.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{client: c}, nil
}

func (c *Client) DebankFeed(ctx context.Context, endBlockCursor *big.Int, accountAddress common.Address) ([]types.Log, error) {
	return nil, nil
}
