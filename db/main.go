package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	url  string
	Pool *pgxpool.Pool
}

func New(url string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &DB{
		Pool: pool,
		url:  url,
	}, nil
}

func (db *DB) CreateSchema(filename string) error {
	schemaSqlBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	schemaSql := string(schemaSqlBytes)

	conn, err := db.Pool.Acquire(context.Background())
	if err != nil {
		return err
	}

	defer conn.Conn().Close(context.Background())

	if _, err := conn.Exec(context.Background(), schemaSql); err != nil {
		return err
	}

	return nil
}
