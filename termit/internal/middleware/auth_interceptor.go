package middleware

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrNoAuthentication = status.Error(codes.Unauthenticated, "authorization is not specified")
	ErrDecodeFailed     = status.Errorf(codes.Unauthenticated, "authorization is not decodable")
	ErrIncorrectFormat  = status.Error(codes.Unauthenticated, "authorization has incorrect format")
	ErrServerFailedAuth = status.Errorf(codes.Internal, "server failed to check authorization")
)

type authInterceptor struct {
	userService *user.Service
}

func NewAuthInterceptor(s *user.Service) *authInterceptor {
	return &authInterceptor{userService: s}
}

func (i *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		u, err := i.authenticate(ctx)
		if err != nil {
			log.Println("Authentication failed:", err)
			return nil, err
		}
		ctx = user.AddToContext(ctx, u.ID)
		return handler(ctx, req)
	}
}

func (i *authInterceptor) authenticate(ctx context.Context) (*user.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNoAuthentication
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, ErrNoAuthentication
	}

	auth, err := decodeBaseAuth(values[0])
	if err != nil {
		return nil, err
	}

	u, err := i.userService.Authenticate(ctx, auth)
	if err == nil { // Successfully authenticated
		return u, nil
	}

	if errors.Is(err, user.ErrNotFound) { // User not found, lets create new one for now
		log.Printf("User with username %s does not exist. Creating...", auth.Username)
		err = i.userService.AddUser(ctx, auth)
		if err != nil {
			return nil, err
		}
		return auth, nil
	}

	return nil, ErrServerFailedAuth
}

func decodeBaseAuth(encoded string) (*user.User, error) {
	if !strings.HasPrefix(encoded, "Basic ") {
		return nil, ErrDecodeFailed
	}
	encoded = encoded[6:]

	dec, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, ErrDecodeFailed
	}

	parts := strings.Split(string(dec), ":")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return nil, ErrIncorrectFormat
	}

	username := parts[0]
	password := parts[1]
	return user.NewUser(username, password), nil
}
