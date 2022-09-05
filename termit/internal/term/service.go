package term

import (
	"context"
	"errors"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

var (
	TermLimitExceeded = errors.New("term limit exceeded")
)

type service struct {
	repo      Repository
	termLimit int
}

func NewService(repo Repository, termLimit int) *service {
	return &service{repo, termLimit}
}

func (s *service) AddTerm(ctx context.Context, t *Term) error {
	atomicAdd := func(r Repository) error {
		cnt, err := r.GetTermCount(ctx, t.Uid)
		if err != nil {
			return err
		}
		if cnt >= s.termLimit {
			return TermLimitExceeded
		}
		return r.AddTerm(ctx, t)
	}
	return s.repo.Atomic(ctx, atomicAdd)
}

func (s *service) GetCollection(ctx context.Context, uid user.UserID) ([]*Term, error) {
	terms, err := s.repo.GetAllTerms(ctx, uid)
	if err != nil {
		return nil, err
	}
	return terms, nil
}
