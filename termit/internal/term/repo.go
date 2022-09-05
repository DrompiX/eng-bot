package term

//go:generate mockgen -destination=repo_mock.go -package=term . Repository

import (
	"context"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

type Repository interface {
	Atomic(context.Context, func(Repository) error) error
	AddTerm(context.Context, *Term) error
	UpdateTerm(context.Context, *Term) error
	DeleteTerm(context.Context, *Term) error
	GetTermCount(context.Context, user.UserID) (int, error)
	GetAllTerms(context.Context, user.UserID) ([]*Term, error)
	GetTermByName(ctx context.Context, uid user.UserID, termName string) (*Term, error)
}
