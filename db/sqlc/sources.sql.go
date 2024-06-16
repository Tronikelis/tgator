// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sources.sql

package sqlc

import (
	"context"
)

const createSource = `-- name: CreateSource :one
INSERT INTO sources (
    id,
    ip 
) VALUES (
    DEFAULT, $1
) RETURNING id, ip
`

func (q *Queries) CreateSource(ctx context.Context, ip string) (Source, error) {
	row := q.db.QueryRow(ctx, createSource, ip)
	var i Source
	err := row.Scan(&i.ID, &i.Ip)
	return i, err
}

const getSourceByIp = `-- name: GetSourceByIp :one
SELECT id, ip FROM sources WHERE ip = $1 LIMIT 1
`

func (q *Queries) GetSourceByIp(ctx context.Context, ip string) (Source, error) {
	row := q.db.QueryRow(ctx, getSourceByIp, ip)
	var i Source
	err := row.Scan(&i.ID, &i.Ip)
	return i, err
}

const getSources = `-- name: GetSources :many
SELECT id, ip FROM sources ORDER BY id DESC
`

func (q *Queries) GetSources(ctx context.Context) ([]Source, error) {
	rows, err := q.db.Query(ctx, getSources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Source
	for rows.Next() {
		var i Source
		if err := rows.Scan(&i.ID, &i.Ip); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
