package db

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	url  string
	Pool *pgxpool.Pool
	PG   goqu.DialectWrapper
}

func New(url string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	goqu.SetDefaultPrepared(true)
	pg := goqu.Dialect("postgres")

	return &DB{
		Pool: pool,
		url:  url,
		PG:   pg,
	}, nil
}

func QueryOne[T any](
	db *DB,
	ctx context.Context,
	query string,
	params ...interface{},
) (T, error) {
	rows, err := db.Pool.Query(ctx, query, params...)
	var t T
	if err != nil {
		return t, err
	}

	return pgx.CollectOneRow(rows, RowToStruct[T])
}

func QueryMany[T any](
	db *DB,
	ctx context.Context,
	query string,
	params ...interface{},
) ([]T, error) {
	rows, err := db.Pool.Query(ctx, query, params...)
	var t []T
	if err != nil {
		return t, err
	}

	return pgx.CollectRows(rows, RowToStruct[T])
}
