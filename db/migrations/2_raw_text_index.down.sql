DROP INDEX messages_raw_text;

DROP EXTENSION pg_trgm;
DROP EXTENSION btree_gin;

CREATE INDEX messages_raw_index ON messages(raw);

