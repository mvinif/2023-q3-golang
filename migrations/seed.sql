CREATE TABLE IF NOT EXISTS person (
			id uuid NOT NULL,
			nickname varchar(32) PRIMARY KEY NOT NULL,
			"name" varchar(100) NOT NULL,
			birthday varchar(10) NULL,
			stacks character varying[] NULL);


-- ALTER TABLE person ADD COLUMN tssearch tsvector GENERATED ALWAYS AS (pg_catalog.to_tsvector('english', name || ' ' || nickname || ' ' || stacks)) STORED;
-- 
-- CREATE TEXT SEARCH CONFIGURATION person_search(COPY=pg_catalog.english);
-- 
-- CREATE OR REPLACE FUNCTION person_search_tsquery(word text) RETURNS tsquery AS 
-- $$
-- BEGIN
--   RETURN plainto_tsquery('public.person_search', trim(word));
-- END
-- $$ LANGUAGE plpgsql;
