package termit

import (
	"context"
	"errors"

	"gitlab.ozon.dev/DrompiX/homework-2/botman/pb"
)

var (
	ServerError = errors.New("server-side error occured")
)

type TermitClient interface {
	AddTerm(context.Context, TermitTerm) error
	GetCollection(context.Context) ([]TermitTerm, error)
	GetTask(context.Context) (TermitTask, error)
	CheckAnswer(context.Context, TermitTaskAnswer) (ValidatedAnswer, error)
}

type GrpcTermitClient struct {
	TermServ pb.TermServiceClient
	TaskServ pb.TaskServiceClient
}

func (c *GrpcTermitClient) AddTerm(ctx context.Context, t TermitTerm) error {
	req := &pb.AddTermRequest{Term: t.Term, Translation: t.Translation}
	_, err := c.TermServ.AddTerm(ctx, req)
	if err != nil {
		err = ServerError
	}
	return err
}

func (c *GrpcTermitClient) GetCollection(ctx context.Context) ([]TermitTerm, error) {
	resp, err := c.TermServ.GetCollection(ctx, &pb.GetCollectionRequest{})
	if err != nil {
		return nil, ServerError
	}

	terms := make([]TermitTerm, 0, len(resp.Terms))
	for _, t := range resp.Terms {
		nt := TermitTerm{Term: t.Term, Translation: t.Translation}
		terms = append(terms, nt)
	}

	return terms, nil
}

func (c *GrpcTermitClient) GetTask(ctx context.Context) (TermitTask, error) {
	resp, err := c.TaskServ.GetTask(ctx, &pb.GetTaskRequest{})
	if err != nil {
		return TermitTask{}, err
	}
	return TermitTask{ID: resp.Id, Term: resp.Term}, nil
}

func (c *GrpcTermitClient) CheckAnswer(ctx context.Context, a TermitTaskAnswer) (ValidatedAnswer, error) {
	req := &pb.CheckAnswerRequest{Id: a.ID, Translation: a.Translation}
	resp, err := c.TaskServ.CheckAnswer(ctx, req)
	if err != nil {
		return ValidatedAnswer{}, err
	}
	return ValidatedAnswer{Success: resp.Success, Answer: resp.Answer, Expect: resp.Expected}, nil
}
