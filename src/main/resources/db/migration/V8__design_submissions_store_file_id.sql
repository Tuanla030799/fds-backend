ALTER TABLE design_submissions
    ADD COLUMN IF NOT EXISTS file_id UUID REFERENCES files(id);

WITH ranked_files AS (
    SELECT id,
           path,
           ROW_NUMBER() OVER (PARTITION BY path ORDER BY created_at DESC, id DESC) AS rn
    FROM files
)
UPDATE design_submissions d
SET file_id = rf.id
FROM ranked_files rf
WHERE d.file_id IS NULL
  AND d.image_url = rf.path
  AND rf.rn = 1;

ALTER TABLE design_submissions
    ALTER COLUMN file_id SET NOT NULL;

ALTER TABLE design_submissions
    DROP COLUMN IF EXISTS image_url;
