-- name: CreateDistrict :one
INSERT INTO district (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetDistrict :one
SELECT * FROM district
WHERE id = $1 LIMIT 1;

-- name: GetAllDistricts :many
SELECT * FROM district
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateDistrict :one
UPDATE district
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteDistrict :exec
DELETE FROM district
WHERE id = $1;
