CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,
    raw text,
    raw_jsonb jsonb,
    created_at timestamp
);
