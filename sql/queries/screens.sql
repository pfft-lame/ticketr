-- name: CreateScreen :one
INSERT INTO screens 
  (name, theater_id, total_seats)
VALUES 
  ( $1, $2, $3 )
RETURNING 
  id, name, total_seats, theater_id;

-- name: UpdateScreenById :one
UPDATE screens SET
  name = COALESCE(sqlc.narg(name), name),
  theater_id = COALESCE(sqlc.narg(theater_id), theater_id),
  total_seats = COALESCE(sqlc.narg(total_seats), total_seats)
WHERE id = $1
RETURNING
  id, name, total_seats, theater_id;

-- name: GetScreenById :one
SELECT id, name, total_seats, theater_id
FROM screens
WHERE id = $1;

-- name: GetAllScreensByTheaterId :many
SELECT id, name, total_seats, theater_id
FROM screens
WHERE theater_id = $1;

-- name: DeleteScreenByID :execrows
DELETE FROM screens
WHERE id = $1;
