-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS goalsCompletions(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	goalId UUID NOT NULL REFERENCES goals(id),
	createdAt timestamptz NOT NULL DEFAULT now()
	);
---- create above / drop below ----
DROP TABLE IRF EXISTS goalsCompletions;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
