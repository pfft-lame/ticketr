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


-- name: GetUpcomingMoviesInTheater :many
SELECT DISTINCT
  m.id,
  m.name,
  m.description,
  m.trailer_url,
  m.languages,
  m.release_date,
  sc.id AS screen_id,
  sc.name AS screen_name
FROM theaters t
JOIN screens sc ON t.id = sc.theater_id
JOIN shows s ON s.screen_id = sc.id
JOIN movies m ON m.id = s.movie_id
JOIN cities c ON c.id = t.city_id
WHERE 
  t.id = $1 AND
  s.start_time > NOW()
