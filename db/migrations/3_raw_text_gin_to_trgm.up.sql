DROP INDEX messages_raw_text;

CREATE INDEX messages_raw_text ON messages USING gin(raw gin_trgm_ops);

