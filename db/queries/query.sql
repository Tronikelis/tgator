-- name: GetMessage :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;
