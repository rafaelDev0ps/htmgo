// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const createSession = `-- name: CreateSession :exec
INSERT INTO sessions (user_id, session_id, expires_at)
VALUES (?, ?, ?)
`

type CreateSessionParams struct {
	UserID    int64
	SessionID string
	ExpiresAt string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.ExecContext(ctx, createSession, arg.UserID, arg.SessionID, arg.ExpiresAt)
	return err
}

const createUser = `-- name: CreateUser :one

INSERT INTO user (email, password, metadata)
VALUES (?, ?, ?)
RETURNING id
`

type CreateUserParams struct {
	Email    string
	Password string
	Metadata interface{}
}

// Queries for User Management
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password, arg.Metadata)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, metadata, created_at, updated_at
FROM user
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password, metadata, created_at, updated_at
FROM user
WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT u.id, u.email, u.password, u.metadata, u.created_at, u.updated_at
FROM user u
         JOIN sessions t ON u.id = t.user_id
WHERE t.session_id = ?
  AND t.expires_at > datetime('now')
`

func (q *Queries) GetUserByToken(ctx context.Context, sessionID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByToken, sessionID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserMetadata = `-- name: UpdateUserMetadata :exec
UPDATE user SET metadata = json_patch(COALESCE(metadata, '{}'), ?) WHERE id = ?
`

type UpdateUserMetadataParams struct {
	JsonPatch interface{}
	ID        int64
}

func (q *Queries) UpdateUserMetadata(ctx context.Context, arg UpdateUserMetadataParams) error {
	_, err := q.db.ExecContext(ctx, updateUserMetadata, arg.JsonPatch, arg.ID)
	return err
}