CREATE TABLE IF NOT EXISTS messages (
  id bigserial PRIMARY KEY,
  raw text,
  created_at timestamp
);
