DROP INDEX IF EXISTS idx_templates_status_sort_created_at;

CREATE INDEX IF NOT EXISTS idx_templates_status_created_at
    ON templates(status, created_at DESC);

ALTER TABLE templates
    DROP COLUMN IF EXISTS tags,
    DROP COLUMN IF EXISTS sort_order;
