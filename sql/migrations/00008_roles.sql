-- +goose Up
CREATE TYPE roles AS ENUM('user', 'admin', 'theater_owner');

CREATE TABLE IF NOT EXISTS user_roles (
  user_id UUID NOT NULL REFERENCES users(id),
  role roles DEFAULT 'user',
  PRIMARY KEY(role, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS user_roles;
DROP TYPE roles;
