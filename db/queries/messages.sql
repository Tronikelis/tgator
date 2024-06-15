-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;

-- name: CreateMessage :exec
INSERT INTO messages (
    id, 
    raw,
    created_at
) VALUES (
    DEFAULT, $1, NOW()
);

-- name: GetMessagesAsc :many
SELECT * FROM messages ORDER BY id ASC LIMIT $1 OFFSET $2;

-- name: GetMessagesDesc :many
SELECT * FROM messages ORDER BY id DESC LIMIT $1 OFFSET $2;

