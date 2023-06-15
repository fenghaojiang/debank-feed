package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type Feed struct {
	TransactionHash common.Hash     `json:"txHash"`
	GasFee          decimal.Decimal `json:"gasFee,omitempty"`
	TokensIncrease  []Token         `json:"tokenIncr,omitempty"`
	TokensDecrease  []Token         `json:"tokenDesc,omitempty"`
}
