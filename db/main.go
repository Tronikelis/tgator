package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

const lockId = 1337

func lockSqlConn(conn *sql.Conn) error {
	if _, err := conn.ExecContext(context.Background(), "SELECT pg_advisory_lock(1)"); err != nil {
		return err
	}

	return nil
}

func unlockSqlConn(conn *sql.Conn) error {
	if _, err := conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock(1)"); err != nil {
		return err
	}

	return nil
}

// concurrency / multi-machine safe migration
func (db *DB) Migrate() error {
	sqlPool, err := sql.Open("postgres", db.url)
	if err != nil {
		return err
	}

	conn, err := sqlPool.Conn(context.Background())
	defer conn.Close()

	if err := lockSqlConn(conn); err != nil {
		sqlPool.Close()
		return err
	}

	driver, err := migrate_postgres.WithInstance(sqlPool, &migrate_postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	defer m.Close()

	defer func() {
		if err := unlockSqlConn(conn); err != nil {
			fmt.Println(err)
			return
		}
	}()

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("no change, continuing")
		return nil
	}
	if err != nil {
		panic(err)
	}

	return nil
}
