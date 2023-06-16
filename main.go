package main

import (
	"context"
	"fmt"

	"github.com/fenghaojiang/debank-feed/client"
	erigonCommon "github.com/ledgerwatch/erigon-lib/common"
)

func main() {
	cli, err := client.NewClient(context.Background(), "https://rpc.ankr.com/eth")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vitalikFeeds, err := cli.DebankFeed(context.Background(), nil, erigonCommon.HexToAddress("0xd8da6bf26964af9d7eed9e03e53415d37aa96045"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, feed := range vitalikFeeds {
		fmt.Printf("%+v", feed)
	}
}
