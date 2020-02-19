-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE posts (
	post_id UUID PRIMARY KEY,
	title VARCHAR (64) NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE posts;
