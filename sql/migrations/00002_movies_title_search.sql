-- +goose Up
-- +goose StatementBegin
ALTER TABLE movies
ADD COLUMN name_tsv tsvector GENERATED ALWAYS AS (
  to_tsvector('english', COALESCE(name, ''))
) STORED;

CREATE INDEX idx_movies_name_tsv
ON movies USING GIN(name_tsv);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE movies
DROP COLUMN name_tsv;
-- +goose StatementEnd
