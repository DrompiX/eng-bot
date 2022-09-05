package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

var userIdKey = "user-id"

var (
	ErrNoUserInMetadata = errors.New("no user id in metadata")
	ErrNoMetadata = errors.New("no metadata for user auth")
)

func AddToContext(ctx context.Context, uid UserID) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.Pairs(userIdKey, string(uid)))
}

func GetFromContext(ctx context.Context) (UserID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrNoMetadata
	}
	if values, exists := md[userIdKey]; exists && len(values) > 0 {
		return UserID(values[0]), nil
	}
	return "", ErrNoUserInMetadata
}