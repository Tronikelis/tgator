-- name: CreateSource :exec
INSERT INTO sources (
    id,
    ip 
) VALUES (
    DEFAULT, $1
);

-- name: GetSources :many
SELECT id, ip FROM sources ORDER BY id DESC;

-- name: GetSourceByIp :one
SELECT id, ip FROM sources WHERE ip = $1 LIMIT 1;
