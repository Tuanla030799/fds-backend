CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_templates_name_trgm
    ON templates USING gin (name gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_templates_note_trgm
    ON templates USING gin (note gin_trgm_ops);
