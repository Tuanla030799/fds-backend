ALTER TABLE templates
    ALTER COLUMN file_id DROP NOT NULL;

ALTER TABLE design_submissions
    ALTER COLUMN file_id DROP NOT NULL;
