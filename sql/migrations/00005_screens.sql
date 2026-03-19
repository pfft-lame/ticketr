-- +goose Up
-- +goose StatementBegin
CREATE TABLE screens (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL,
  theater_id UUID NOT NULL REFERENCES theaters(id) ON DELETE CASCADE,
  total_seats INTEGER NOT NULL,
  created_at timestamptz DEFAULT NOW(),
  updated_at timestamptz DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_theater_id_name
ON screens(LOWER(name), theater_id);

CREATE INDEX idx_theater_id
ON screens(theater_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE screens;
-- +goose StatementEnd
