-- name: CreateShow :one
INSERT INTO shows 
  ( movie_id, screen_id, start_time, end_time )
VALUES 
  ( $1, $2, $3, $4 )
RETURNING 
  id, movie_id, screen_id, start_time, end_time;


-- name: UpdateShowById :one
UPDATE shows SET 
  movie_id = COALESCE(sqlc.narg(movie_id), movie_id),
  screen_id = COALESCE(sqlc.narg(screen_id), screen_id),
  start_time = COALESCE(sqlc.narg(start_time), start_time),
  end_time = COALESCE(sqlc.narg(end_time), end_time)
WHERE 
  id = $1
RETURNING
  id, movie_id, screen_id, start_time, end_time;


-- name: DeleteShowById :execrows
DELETE FROM shows
WHERE id = $1;

-- name: GetShowInfoById :one
SELECT 
  s.id,
  s.start_time,
  s.end_time,
  s.screen_id,
  s.movie_id,
  m.name AS movie_name,
  sc.name AS screen_name,
  t.id AS theater_id,
  t.name AS theater_name
FROM shows s
JOIN movies m ON m.id = s.movie_id
JOIN screens sc ON sc.id = s.screen_id
JOIN theaters t ON t.id = sc.theater_id
WHERE s.id = $1;


-- name: GetShowsByMovieId :many
SELECT
  s.id,
  s.start_time,
  s.end_time,
  s.screen_id,
  sc.name AS screen_name,
  t.id AS theater_id,
  t.name AS theater_name
FROM movies m
JOIN shows s ON s.movie_id = m.id
JOIN screens sc ON sc.id = s.screen_id
JOIN theaters t ON t.id = sc.theater_id
WHERE 
  m.id = sqlc.arg(movie_id) AND
  s.start_time > NOW();

-- name: GetShowsByTheaterId :many
SELECT 
  m.id AS movie_id,
  m.name AS movie_name,
  sc.id AS screen_id,
  sc.name AS screen_name,
  s.start_time,
  s.end_time
FROM shows s
JOIN movies m ON m.id = s.movie_id
JOIN screens sc ON sc.id = s.screen_id
JOIN theaters t ON t.id = sc.theater_id
WHERE 
  t.id = sqlc.arg(theater_id) AND
  s.screen_id > NOW();

-- name: GetShowsBetweenTimeRange :many
