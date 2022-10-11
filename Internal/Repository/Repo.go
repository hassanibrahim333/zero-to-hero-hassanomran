package Repository

import (
	"context"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
	"time"
)

type Transactions struct {
	bun.BaseModel   `bun:"table:transaction"`
	ID              string          `bun:"id,pk,notnull"`
	Amount          decimal.Decimal `bun:"amount"`
	Currency        string          `bun:"currency"`
	CreatedAt       time.Time       `bun:"createdat"`
	StatusId        int8            `bun:"statusid"`
	AuthorizationId string          `bun:"authorizationid"`
}
type databaseRepo struct {
	conn *bun.DB
}

func NewDefaultRepository(conn *bun.DB) *databaseRepo {
	return &databaseRepo{
		conn: conn,
	}
}

func (db *databaseRepo) List(ctx context.Context) ([]Transactions, error) {
	var models []Transactions
	err := db.conn.NewSelect().Model(&models).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (db *databaseRepo) Create(ctx context.Context, model *Transactions) error {
	_, err := db.conn.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (db *databaseRepo) Update(ctx context.Context, model *Transactions, AuthorizationId string) error {
	_, err := db.conn.NewUpdate().Model(model).Column("statusid").Where("authorizationid = ?", AuthorizationId).Exec(ctx)
	if err != nil {
		return nil
	}
	return nil
}
