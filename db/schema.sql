CREATE TABLE IF NOT EXISTS sources (
    id serial PRIMARY KEY,
    ip varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,
    raw text,
    raw_jsonb jsonb,
    created_at timestamp NOT NULL,

    source_id int REFERENCES sources ON DELETE CASCADE NOT NULL
);

