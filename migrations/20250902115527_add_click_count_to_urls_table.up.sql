BEGIN;

ALTER TABLE urls
ADD COLUMN click_count INTEGER
DEFAULT 0 CHECK (click_count >= 0);

COMMIT;

