package Models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID              string          `json:"id" validate:"required`
	Amount          decimal.Decimal `json:"Amount" validate:"required`
	Currency        string          `json:"Currency" validate:"required,len=3`
	CreatedAt       time.Time       `json:"CreatedAt" validate:"required`
	StatusId        int8            `json:"StatusId"`
	AuthorizationId string          `json:"AuthorizationId"`
}
