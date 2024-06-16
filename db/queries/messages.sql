-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;

-- name: CreateMessage :one
INSERT INTO messages (
    id, 
    created_at,
    raw,
    source_id
) VALUES (
    DEFAULT, NOW(), $1, $2
) RETURNING *;

-- name: GetMessagesAsc :many
SELECT id, created_at, raw FROM messages ORDER BY id ASC LIMIT $1 OFFSET $2;

-- name: GetMessagesDesc :many
SELECT id, created_at, raw FROM messages ORDER BY id DESC LIMIT $1 OFFSET $2;

-- name: GetMessagesWhereSourceId :many
SELECT id, raw, created_at FROM messages WHERE source_id = $1 AND raw ilike $2
    ORDER BY id DESC LIMIT $3 OFFSET $4;
