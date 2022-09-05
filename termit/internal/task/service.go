package task

import (
	"context"
	"errors"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/term"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

var (
	ErrAlreadyAnswered = errors.New("task already completed")
)

type Service struct {
	taskRepo Repository
	termRepo term.Repository
}

func NewService(taskRepo Repository, termRepo term.Repository) *Service {
	return &Service{taskRepo: taskRepo, termRepo: termRepo}
}

func (s *Service) GenerateTask(ctx context.Context, uid user.UserID) (*task, error) {
	terms, err := s.termRepo.GetAllTerms(ctx, uid)
	if err != nil {
		return nil, err
	}
	if len(terms) == 0 {
		return nil, ErrCollectionIsEmpty
	}
	randTerm := terms[GetRandomNonNegInt(len(terms))]

	t := NewTask(randTerm.Uid, randTerm.Data, randTerm.Translation)
	if err = s.taskRepo.Create(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) CheckAnswer(ctx context.Context, a *answer) (*AnswerCheck, error) {
	task, err := s.taskRepo.GetById(ctx, a.Tid)
	if err != nil {
		return nil, err
	}

	if task.Success != nil {
		return nil, ErrAlreadyAnswered
	}

	if task.Uid != a.Uid {
		return nil, ErrNotFound
	}

	success := a.Translation == task.Expected
	task.Success = &success
	
	if err = s.taskRepo.Update(ctx, &task); err != nil {
		return nil, err
	}

	return &AnswerCheck{Success: success, Expected: task.Expected}, nil
}
