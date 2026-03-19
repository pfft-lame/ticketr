-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS theaters (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(70) NOT NULL,
  description TEXT NOT NULL,
  city_id UUID NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
  address TEXT NOT NULL,
  pincode CHAR(6) NOT NULL CHECK ( pincode ~ '^[0-9]{6}$' ),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS theaters;
-- +goose StatementEnd
