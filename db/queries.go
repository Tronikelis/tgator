package db

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5"
)

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

func CountRows(db *DB, ctx context.Context, sd *goqu.SelectDataset) (int32, error) {
	query, params, err := sd.
		Select(goqu.L("count(*)")).
		ClearOrder().
		ClearOffset().
		ClearLimit().
		ToSQL()

	if err != nil {
		return 0, err
	}

	count, err := QueryOne[int32](db, ctx, query, params...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
