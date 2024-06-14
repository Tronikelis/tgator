-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;

-- name: CreateMessage :exec
INSERT INTO messages (
  id, raw, created_at
) VALUES (
  DEFAULT, $1, NOW()
);
