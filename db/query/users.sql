-- name: CreateUser :one
INSERT INTO users (
    username, hashed_password, full_name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    hashed_password = $3,
    full_name = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;