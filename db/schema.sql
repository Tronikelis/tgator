CREATE TABLE IF NOT EXISTS sources (
    id serial PRIMARY KEY,
    ip varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,
    raw text,
    created_at timestamp NOT NULL DEFAULT NOW(),

    source_id int REFERENCES sources ON DELETE CASCADE NOT NULL
);

CREATE INDEX messages_raw_index ON messages(raw);
