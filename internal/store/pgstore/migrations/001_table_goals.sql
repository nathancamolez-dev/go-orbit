-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS goals (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	title text NOT NULL,
	desiredWeeklyFrequency int NOT NULL,
	createdAt timestamptz NOT NULL DEFAULT now()
);
---- create above / drop below ----
DROP TABALE IF EXISTS goals;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
