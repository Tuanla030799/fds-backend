ALTER TABLE presets
    ADD COLUMN IF NOT EXISTS file_id UUID REFERENCES files(id);

WITH ranked_files AS (
    SELECT id,
           path,
           ROW_NUMBER() OVER (PARTITION BY path ORDER BY created_at DESC, id DESC) AS rn
    FROM files
)
UPDATE presets p
SET file_id = rf.id
FROM ranked_files rf
WHERE p.file_id IS NULL
  AND p.image_url = rf.path
  AND rf.rn = 1;

ALTER TABLE presets
    ALTER COLUMN file_id SET NOT NULL;

ALTER TABLE presets
    DROP COLUMN IF EXISTS image_url;
