package task

//go:generate mockgen -destination=repo_mock.go -package=task . Repository

import (
	"context"
)

type Repository interface {
	Create(context.Context, *task) error
	Update(context.Context, *task) error
	GetById(context.Context, TaskID) (task, error)
}
