-- name: CreateSource :exec
INSERT INTO sources (
    id,
    ip
) VALUES (
    DEFAULT, $1
);

-- name: GetSources :many
SELECT id, ip FROM sources ORDER BY id DESC;

