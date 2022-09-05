package user

import (
	"context"
	"crypto/sha512"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrIncorrectPassword = status.Error(codes.Unauthenticated, "incorrect password")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Authenticate(ctx context.Context, u *User) (*User, error) {
	dbUser, err := s.repo.Find(ctx, u.Username)
	if err != nil {
		return nil, err
	}
	if err := compareHashToAuthPass(dbUser.Password, u.Password); err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (s *Service) AddUser(ctx context.Context, u *User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return s.repo.Save(ctx, u)
}

// NOTE: Hashing with bcrypt or SHA can be hidden behind a shared interface
// and injected in service constructor, but now don't have enough time for this refactoring

func hashPassword(p string) (string, error) {
	// TODO: go back to bcrypt implementation when sessions/JWT will be added
	// password, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	// if err != nil {
	// 	return "", err
	// }
	// return string(password), nil

	hash := sha512.Sum512([]byte(p))
	return hex.EncodeToString(hash[:]), nil
}

func compareHashToAuthPass(passHash, authPass string) error {
	// NOTE: Takes 1+ second to compare with bcrypt not tolerable for auth on each request
	// if err := bcrypt.CompareHashAndPassword(passHash, authPass); err != nil {
	// 	return ErrIncorrectPassword
	// }
	// return nil

	authPassHash, err := hashPassword(authPass)
	if err != nil {
		return err
	}
	if passHash != authPassHash {
		return ErrIncorrectPassword
	}
	return nil
}
