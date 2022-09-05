package user

//go:generate mockgen -destination=repo_mock.go -package=user . Repository

import (
	"context"
	"errors"
)

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNotFound      = errors.New("user not found")
)

type Repository interface {
	Save(context.Context, *User) error
	Find(ctx context.Context, username string) (*User, error)
}
