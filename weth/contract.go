package weth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fenghaojiang/debank-feed/constant"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./weth.abi --pkg weth --type weth --out ./weth.go

func FilterWETHDeposit(l types.Log) bool {
	return l.Topics[0] == constant.DepositEvent
}

func FilterWETHWithdrawal(l types.Log) bool {
	return l.Topics[0] == constant.WithdrawalEvent
}
