-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetAllUsers :many
SELECT * FROM users;


-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
