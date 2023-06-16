package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fenghaojiang/debank-feed/erc20"
	"github.com/fenghaojiang/debank-feed/erc721"
	"github.com/fenghaojiang/debank-feed/weth"
	erigonCommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/eth/filters"
)

type Client struct {
	client       *rpc.Client
	Erc20Filter  *erc20.ERC20Filterer
	Erc721Filter *erc721.ERC721Filterer
	WethFilter   *weth.WethFilterer
}

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	c, err := rpc.DialContext(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	erc20filter, _ := erc20.NewERC20Filterer(erigonCommon.HexToAddress(""), nil)
	erc721filter, _ := erc721.NewERC721Filterer(erigonCommon.HexToAddress(""), nil)
	wethfilter, _ := weth.NewWethFilterer(erigonCommon.HexToAddress(""), nil)

	return &Client{
		client:       c,
		Erc20Filter:  erc20filter,
		Erc721Filter: erc721filter,
		WethFilter:   wethfilter,
	}, nil
}

func (c *Client) DebankFeed(ctx context.Context, endBlockCursor *big.Int, accountAddress erigonCommon.Address) ([]types.Log, error) {
	var logs []types.Log

	query := filters.FilterCriteria{
		FromBlock: new(big.Int).SetUint64(0),
		Topics: [][]erigonCommon.Hash{
			{
				accountAddress.Hash(),
			},
		},
	}

	if endBlockCursor != nil {
		query.ToBlock = endBlockCursor
	}

	err := c.client.CallContext(ctx, &logs, "erigon_getLatestLogs", query, filters.LogFilterOptions{
		LogCount:          20,
		IgnoreTopicsOrder: true,
	})
	return logs, err
}
