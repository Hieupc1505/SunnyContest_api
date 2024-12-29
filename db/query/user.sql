-- name: GetUserByEmail :one
SELECT id, username, password, role, status, token, token_expried, created_time, updated_time
FROM sf_user
WHERE username = $1;

-- name: CreateUser :one
INSERT INTO sf_user (username, password, role, status, token, token_expried, created_time, updated_time)
VALUES ($1, $2, $3, $4, $5, $6, now(), now())
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time;

-- name: ChangePassword :one
UPDATE sf_user
SET password = $1, updated_time = now()
WHERE id = $2
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time;

-- name: DeleteUser :exec
DELETE FROM sf_user
WHERE id = $1;

-- name: GetUserByID :one
SELECT id, username, password, role, status, token, token_expried, created_time, updated_time
FROM sf_user
WHERE id = $1;

-- name: UpdateUser :one
UPDATE sf_user
SET username = $1, password = $2, role = $3, status = $4, token = $5, token_expried = $6, updated_time = now()
WHERE id = $7
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time;
