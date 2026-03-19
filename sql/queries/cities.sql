-- name: CreateCity :one
INSERT INTO cities ( city, state)
VALUES ( $1, $2 )
RETURNING id, city, state;


-- name: DeleteCityById :execrows
DELETE FROM cities
WHERE id = $1;


-- name: GetAllCities :many
SELECT id, city, state FROM cities;


-- name: GetCityById :one
SELECT city, state FROM cities
WHERE id = $1;
