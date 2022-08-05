-- name: CreateSchool :one
INSERT INTO school (
    name, district_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSchool :one
SELECT * FROM school
WHERE id = $1 LIMIT 1;

-- name: GetAllSchools :many
SELECT * FROM school
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateSchool :one
UPDATE school
SET name = $2,
    district_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSchool :exec
DELETE FROM school
WHERE id = $1;