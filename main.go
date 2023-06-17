package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/debank-feed/client"
	"github.com/fenghaojiang/debank-feed/erc20"
	"github.com/fenghaojiang/debank-feed/erc721"
)

func main() {
	cli, err := client.NewClient(context.Background(), "http://localhost:8545")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	account := common.HexToAddress("0xd8da6bf26964af9d7eed9e03e53415d37aa96045")

	vitalikLogs, err := cli.DebankFeed(context.Background(), nil, account)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("length of vitalik logs: %d", len(vitalikLogs))

	// token balance changes in the same transaction
	tokenChanges := make(map[common.Hash]map[common.Address]*big.Int)

	for _, l := range vitalikLogs {
		switch {
		case erc20.FilterERC20Transfer(l):
			if tokenChanges[l.TxHash] == nil {
				tokenChanges[l.TxHash] = make(map[common.Address]*big.Int)
			}

			event, err := cli.Erc20Filter.ParseTransfer(l)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if tokenChanges[l.TxHash][l.Address] == nil {
				tokenChanges[l.TxHash][l.Address] = new(big.Int).SetUint64(0)
			}

			// calculate the aggregated result
			if event.From == account {
				tokenChanges[l.TxHash][l.Address].Sub(tokenChanges[l.TxHash][l.Address], event.Value)
			}

			if event.To == account {
				tokenChanges[l.TxHash][l.Address].Add(tokenChanges[l.TxHash][l.Address], event.Value)
			}

		case erc721.FilterERC721Transfer(l):
			if tokenChanges[l.TxHash] == nil {
				tokenChanges[l.TxHash] = make(map[common.Address]*big.Int)
			}

			event, err := cli.Erc721Filter.ParseTransfer(l)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if tokenChanges[l.TxHash][l.Address] == nil {
				tokenChanges[l.TxHash][l.Address] = new(big.Int).SetUint64(0)
			}

			// calculate the aggregated result
			if event.From == account {
				tokenChanges[l.TxHash][l.Address].Sub(tokenChanges[l.TxHash][l.Address], new(big.Int).SetUint64(1))
			}

			if event.To == account {
				tokenChanges[l.TxHash][l.Address].Add(tokenChanges[l.TxHash][l.Address], new(big.Int).SetUint64(1))
			}

			// TODO Weth Aggregated Logic
		}
	}
}
