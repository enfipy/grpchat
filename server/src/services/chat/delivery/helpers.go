package delivery

import (
	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func getUsernameFromContext(ctx context.Context) string {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		panic(errors.New("Can not get metadata"))
	}

	for key, value := range headers {
		if key == "username" {
			length := len(value)
			if length < 1 || length > 1 {
				panic(errors.New("Invalid metadata: username"))
			}
			return value[0]
		}
	}

	return ""
}
