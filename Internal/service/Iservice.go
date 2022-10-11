package service

import (
	"context"
	"secondchallange/Internal/Models"
)

type IService interface {
	List(ctx context.Context) ([]Models.Transaction, error)
	Create(ctx context.Context, transaction *Models.Transaction) (*Models.Transaction, error)
}
