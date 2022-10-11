package Repository

import (
	"context"
)

type IRepo interface {
	List(ctx context.Context) ([]Transactions, error)
	Create(ctx context.Context, transactions *Transactions) error
	Update(ctx context.Context, transactions *Transactions, AuthorizationId string) error
}
