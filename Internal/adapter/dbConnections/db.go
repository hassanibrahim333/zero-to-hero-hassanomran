package dbConnections

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"secondchallange/config"
)

func ConnectToDatabase(config config.Configurations) (*bun.DB, error) {
	connectionString := config.Database.ConnectionString
	ctx := context.Background()
	con, err := pgx.Connect(ctx, connectionString)
	defer con.Close(context.Background())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	return db, nil
}
