package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fenghaojiang/debank-feed/client"
)

func main() {
	cli, err := client.NewClient(context.Background(), "http://localhost:8545")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vitalikLogs, err := cli.DebankFeed(context.Background(), nil, common.HexToAddress("0xd8da6bf26964af9d7eed9e03e53415d37aa96045"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, feed := range vitalikLogs {
		switch feed.Topics[0] {

		}
	}
}
