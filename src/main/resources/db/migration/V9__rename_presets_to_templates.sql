ALTER TABLE presets RENAME TO templates;

ALTER INDEX IF EXISTS idx_presets_status_sort_created_at
    RENAME TO idx_templates_status_sort_created_at;

ALTER TABLE templates
    ALTER COLUMN status SET DEFAULT 'ACTIVE';

UPDATE templates
SET status = UPPER(status)
WHERE status IS NOT NULL;

ALTER TABLE templates
    DROP CONSTRAINT IF EXISTS chk_templates_status;

ALTER TABLE templates
    ADD CONSTRAINT chk_templates_status
    CHECK (status IN ('ACTIVE', 'INACTIVE'));

DROP TRIGGER IF EXISTS trg_presets_set_updated_at ON templates;
DROP TRIGGER IF EXISTS trg_templates_set_updated_at ON templates;

CREATE TRIGGER trg_templates_set_updated_at
BEFORE UPDATE ON templates
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
