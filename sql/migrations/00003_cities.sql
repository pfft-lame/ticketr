-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cities (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  city VARCHAR(160) NOT NULL,
  state VARCHAR(160) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_city_state
ON cities(LOWER(city), LOWER(state));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cities;
-- +goose StatementEnd
