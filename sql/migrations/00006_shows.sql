-- +goose Up
-- +goose StatementBegin
CREATE TABLE shows (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  movie_id UUID NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
  screen_id UUID NOT NULL REFERENCES screens(id) ON DELETE CASCADE,
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT valid_time CHECK ( end_time > start_time )
);

CREATE EXTENSION IF NOT EXISTS btree_gist;

ALTER TABLE shows
ADD CONSTRAINT no_overlap
EXCLUDE USING gist (
  screen_id WITH =,
  tstzrange(start_time, end_time, '[)') with &&
);
-- [ -> include start time ---- ) -> exclude start time 
-- also we don't necessarily have to pass this parameter because the default behaviour is the same
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shows;
-- +goose StatementEnd
