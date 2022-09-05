package task

import (
	"context"
	"errors"
	"log"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type taskServer struct {
	service *Service
	pb.UnimplementedTaskServiceServer
}

func NewServer(s *Service) *taskServer {
	return &taskServer{service: s}
}

func (s *taskServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskReponse, error) {
	userId, err := user.GetFromContext(ctx)
	if err != nil {
		log.Printf("Unexpected unauthorized request reached server: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized request")
	}

	genTask, err := s.service.GenerateTask(ctx, userId)
	if err != nil {
		if errors.Is(err, ErrCollectionIsEmpty) {
			err = status.Error(codes.NotFound, err.Error())
		} else {
			err = status.Errorf(codes.Internal, "could not generate new task: %s", err)
		}
		return nil, err
	}

	grpcTask := &pb.GetTaskReponse{
		Id: string(genTask.ID),
		Term: genTask.Term,
	}

	return grpcTask, nil
}

func (s *taskServer) CheckAnswer(ctx context.Context, req *pb.CheckAnswerRequest) (*pb.CheckAnswerResponse, error) {
	userId, err := user.GetFromContext(ctx)
	if err != nil {
		log.Printf("Unexpected unauthorized request reached server: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized request")
	}
	ans := NewAnswer(TaskID(req.Id), userId, req.Translation)
	check, err := s.service.CheckAnswer(ctx, ans)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check answer: %s", err)
	}

	resp := &pb.CheckAnswerResponse{
		Success: check.Success,
		Answer: ans.Translation,
		Expected: check.Expected,
	}

	return resp, nil
}
