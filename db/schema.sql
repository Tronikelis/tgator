CREATE TABLE IF NOT EXISTS messages (
  id serial PRIMARY KEY,
  raw text,
  created_at timestamp
);
