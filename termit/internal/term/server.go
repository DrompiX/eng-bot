package term

import (
	"context"
	"errors"
	"log"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/pb"
)

type termServer struct {
	service *service
	pb.UnimplementedTermServiceServer
}

func NewServer(s *service) *termServer {
	return &termServer{service: s}
}

func (s *termServer) AddTerm(ctx context.Context, addReq *pb.AddTermRequest) (*pb.AddTermResponse, error) {
	userId, err := user.GetFromContext(ctx)
	if err != nil {
		log.Printf("Unexpected unauthorized request reached server: %s", err)
		return nil, errors.New("unauthorized request")
	}

	t, err := NewTerm(addReq.Term, addReq.Translation, userId)
	if err != nil {
		return nil, err
	}
	if err = s.service.AddTerm(ctx, t); err != nil {
		return nil, err
	}
	return &pb.AddTermResponse{}, nil
}

func (s *termServer) GetCollection(ctx context.Context, req *pb.GetCollectionRequest) (*pb.GetCollectionResponse, error) {
	userId, err := user.GetFromContext(ctx)
	if err != nil {
		log.Printf("Unexpected unauthorized request reached server: %s", err)
		return nil, errors.New("unauthorized request")
	}

	col, err := s.service.GetCollection(ctx, userId)
	if err != nil {
		return nil, err
	}

	terms := make([]*pb.GetCollectionResponse_TermInfo, 0, len(col))
	for i := 0; i < len(col); i++ {
		t := &pb.GetCollectionResponse_TermInfo{
			Term: col[i].Data,
			Translation: col[i].Translation,
		}
		terms = append(terms, t)
	}
	res := &pb.GetCollectionResponse{Terms: terms}
	return res, nil
}
