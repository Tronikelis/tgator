DROP INDEX messages_raw_text;

CREATE INDEX CREATE INDEX messages_raw_text ON messages USING gin(raw);


