-- name: CreateTheater :one
INSERT INTO theaters (
  name, description, city_id, address, pincode
) VALUES ($1, $2, $3, $4, $5)
RETURNING 
  id, name, description, city_id, address, pincode;


-- name: DeleteTheaterById :execrows
DELETE FROM theaters
WHERE id = $1;


-- name: UpdateTheatreById :one
UPDATE theaters SET
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description),
  city_id = COALESCE(sqlc.narg(city_id), city_id),
  address = COALESCE(sqlc.narg(address), address),
  pincode = COALESCE(sqlc.narg(pincode), pincode),
  updated_at = NOW()
WHERE id = $1
RETURNING 
  id, name, description, city_id, address, pincode;


-- name: GetAllTheaters :many
SELECT 
  id,
  name,
  description,
  city_id,
  address,
  pincode
FROM theaters;


-- name: GetTheatersById :one
SELECT 
  id,
  name,
  description,
  city_id,
  address,
  pincode
FROM theaters
WHERE id = $1;


-- name: GetTheatersByCityId :many
SELECT 
  id,
  name,
  description,
  city_id,
  address,
  pincode
FROM theaters
WHERE city_id = $1;
