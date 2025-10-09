// middleware/auth_interceptor.go
package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata not provided")
		}

		userIDs := md.Get("user_id")
		if len(userIDs) == 0 {
			return nil, status.Error(codes.Unauthenticated, "user_id not provided")
		}

		// Добавляем userID в контекст
		ctx = context.WithValue(ctx, userIDKey, userIDs[0])
		return handler(ctx, req)
	}
}

// Хелпер для извлечения userID из контекста
func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok || userID == "" {
		return "", status.Error(codes.Unauthenticated, "user not authenticated")
	}
	return userID, nil
}
