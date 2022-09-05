package termit

import (
	"context"
	"encoding/base64"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func PrepareAuthCtx(ctx context.Context, chatId int64) context.Context {
	auth := getBasicAuth(chatId)
	return metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
}

// NOTE: Not secure at all, was done just for fast testing
func getBasicAuth(chatId int64) string {
	username := chatId
	password := chatId * 42 + 579
	creds := fmt.Sprintf("%d:%d", username, password)
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))
}