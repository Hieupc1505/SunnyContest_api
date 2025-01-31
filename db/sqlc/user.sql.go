// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const changePassword = `-- name: ChangePassword :one
UPDATE sf_user
SET password = $1, updated_time = now()
WHERE id = $2
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time
`

type ChangePasswordParams struct {
	Password string `json:"password"`
	ID       int64  `json:"id"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) (SfUser, error) {
	row := q.db.QueryRow(ctx, changePassword, arg.Password, arg.ID)
	var i SfUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.TokenExpried,
		&i.CreatedTime,
		&i.UpdatedTime,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO sf_user (username, password, role, status, token, token_expried, created_time, updated_time)
VALUES ($1, $2, $3, $4, $5, $6, now(), now())
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time
`

type CreateUserParams struct {
	Username     string      `json:"username"`
	Password     string      `json:"password"`
	Role         int32       `json:"role"`
	Status       int32       `json:"status"`
	Token        pgtype.Text `json:"token"`
	TokenExpried pgtype.Int8 `json:"token_expried"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (SfUser, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Role,
		arg.Status,
		arg.Token,
		arg.TokenExpried,
	)
	var i SfUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.TokenExpried,
		&i.CreatedTime,
		&i.UpdatedTime,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM sf_user
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, password, role, status, token, token_expried, created_time, updated_time
FROM sf_user
WHERE username = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, username string) (SfUser, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, username)
	var i SfUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.TokenExpried,
		&i.CreatedTime,
		&i.UpdatedTime,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password, role, status, token, token_expried, created_time, updated_time
FROM sf_user
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (SfUser, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i SfUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.TokenExpried,
		&i.CreatedTime,
		&i.UpdatedTime,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE sf_user
SET username = $1, password = $2, role = $3, status = $4, token = $5, token_expried = $6, updated_time = now()
WHERE id = $7
    RETURNING id, username, password, role, status, token, token_expried, created_time, updated_time
`

type UpdateUserParams struct {
	Username     string      `json:"username"`
	Password     string      `json:"password"`
	Role         int32       `json:"role"`
	Status       int32       `json:"status"`
	Token        pgtype.Text `json:"token"`
	TokenExpried pgtype.Int8 `json:"token_expried"`
	ID           int64       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (SfUser, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.Password,
		arg.Role,
		arg.Status,
		arg.Token,
		arg.TokenExpried,
		arg.ID,
	)
	var i SfUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.TokenExpried,
		&i.CreatedTime,
		&i.UpdatedTime,
	)
	return i, err
}
