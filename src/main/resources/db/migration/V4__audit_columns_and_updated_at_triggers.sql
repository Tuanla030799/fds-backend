ALTER TABLE files
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    ADD COLUMN IF NOT EXISTS created_by UUID REFERENCES admin_users(id),
    ADD COLUMN IF NOT EXISTS updated_by UUID REFERENCES admin_users(id);

ALTER TABLE presets
    ADD COLUMN IF NOT EXISTS created_by UUID REFERENCES admin_users(id),
    ADD COLUMN IF NOT EXISTS updated_by UUID REFERENCES admin_users(id);

ALTER TABLE design_submissions
    ADD COLUMN IF NOT EXISTS created_by UUID REFERENCES admin_users(id),
    ADD COLUMN IF NOT EXISTS updated_by UUID REFERENCES admin_users(id);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_presets_set_updated_at ON presets;
CREATE TRIGGER trg_presets_set_updated_at
BEFORE UPDATE ON presets
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

DROP TRIGGER IF EXISTS trg_design_submissions_set_updated_at ON design_submissions;
CREATE TRIGGER trg_design_submissions_set_updated_at
BEFORE UPDATE ON design_submissions
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

DROP TRIGGER IF EXISTS trg_files_set_updated_at ON files;
CREATE TRIGGER trg_files_set_updated_at
BEFORE UPDATE ON files
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
