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
RETURNING id;

-- name: GetMovieById :one
SELECT
  id, name, description, casts, trailer_url, languages, release_date, director, status
FROM movies
WHERE id = $1;

-- name: DeleteMovieById :exec
DELETE FROM movies
WHERE id = $1;

-- name: UpdateMovieById :one
UPDATE movies SET 
  name = coalesce(sqlc.narg(name), name),
  description = coalesce(sqlc.narg(description), description),
  casts = coalesce(sqlc.narg(casts), casts),
  trailer_url = coalesce(sqlc.narg(trailer_url), trailer_url),
  languages = coalesce(sqlc.narg(languages), languages),
  release_date = coalesce(sqlc.narg(release_date), release_date),
  director = coalesce(sqlc.narg(director), director),
  status = coalesce(sqlc.narg(status), status),
  updated_at = NOW()
WHERE id = $1
RETURNING
  name, description, casts, trailer_url, languages, release_date, director, status;

-- name: GetAllMovies :many
SELECT 
  name, description, casts, trailer_url, languages, release_date, director, status
FROM movies;

-- ame: GetMovieByName :many
-- SELECT 
--   name, description, casts, trailer_url, languages, release_date, director, status
-- FROM movies
-- WHERE name_tsv @@ websearch_to_tsquery('english', $1)
-- ORDER BY rank DESC;
