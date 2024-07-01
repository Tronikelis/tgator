DROP INDEX messages_raw_index;

CREATE EXTENSION pg_trgm;
CREATE EXTENSION btree_gin;

CREATE INDEX messages_raw_text ON messages USING gin(raw);




