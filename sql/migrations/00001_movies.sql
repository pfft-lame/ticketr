-- +goose Up
-- +goose StatementBegin
CREATE TYPE release_status AS ENUM('RELEASED', 'UNRELEASED', 'BLOCKED');

CREATE TABLE IF NOT EXISTS movies (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(70) NOT NULL,
  description TEXT NOT NULL,
  casts TEXT[] NOT NULL,
  trailer_url TEXT NOT NULL,
  languages TEXT[] NOT NULL,
  release_date TIMESTAMPTZ NOT NULL,
  director VARCHAR(70) NOT NULL,
  status RELEASE_STATUS NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd
--
--
--
--
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
DROP TYPE release_status;
-- +goose StatementEnd
