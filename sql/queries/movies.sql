-- name: CreateMovie :one
INSERT INTO
  movies (
    name,
    description,
    casts,
    trailer_url,
    languages,
    release_date,
    director,
    status
  )
VALUES
  ( $1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING 
  id, name, description, casts, trailer_url, languages, release_date, director, status;


-- name: GetMovieById :one
SELECT
  id, name, description, casts, trailer_url, languages, release_date, director, status
FROM movies
WHERE id = $1;


-- name: DeleteMovieById :execrows
DELETE FROM movies
WHERE id = $1;


-- name: UpdateMovieById :one
UPDATE movies SET 
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description),
  casts = COALESCE(sqlc.narg(casts), casts),
  trailer_url = COALESCE(sqlc.narg(trailer_url), trailer_url),
  languages = COALESCE(sqlc.narg(languages), languages),
  release_date = COALESCE(sqlc.narg(release_date), release_date),
  director = COALESCE(sqlc.narg(director), director),
  status = COALESCE(sqlc.narg(status), status),
  updated_at = NOW()
WHERE id = $1
RETURNING
  name, description, casts, trailer_url, languages, release_date, director, status;


-- name: GetAllMovies :many
SELECT 
  name, description, casts, trailer_url, languages, release_date, director, status
FROM movies;

-- name: GetUpcomingMovies :many
SELECT 
  name, description, casts, trailer_url, languages, release_date, director, status
FROM movies
WHERE release_date > NOW();

-- name: GetAiringMovies :many
SELECT 
  m.name, m.description, casts, trailer_url, languages, release_date, director, status
FROM movies m
JOIN shows s on s.movie_id = m.id
JOIN screens sc on sc.id = s.screen_id
JOIN theaters t ON t.id = sc.theater_id
JOIN cities c ON c.id = t.city_id
WHERE
  c.id = sqlc.arg(city_id) AND
  s.start_time >= NOW() AND
  s.end_time < (CURRENT_DATE + INTERVAL '1 day');


-- -- name: GetMovieByName :many
SELECT 
  id, name, description, casts, trailer_url, languages, release_date, director, status,
  ts_rank(name_tsv, websearch_to_tsquery('english', $1)) AS rank
FROM movies
WHERE name_tsv @@ websearch_to_tsquery('english', $1)
ORDER BY rank DESC;
