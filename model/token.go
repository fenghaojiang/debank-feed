package model

import (
	"github.com/shopspring/decimal"
)

type Token struct {
	Name    string          `json:"name"`
	Symbol  string          `json:"symbol"`
	Decimal uint8           `json:"decimal"`
	Value   decimal.Decimal `json:"value"`
}
