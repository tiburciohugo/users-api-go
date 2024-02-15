-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindUserByEmail :one
SELECT u.id, u.name, u.email FROM users u WHERE u.email = $1;

