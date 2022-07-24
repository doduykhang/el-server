package util

import (
	"context"
	"net/http"
)

var (
	userIdKey = "user_id"
)

func ContextWithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIdKey, userID)
}

func RequestWithUserID(r *http.Request, userID int64) *http.Request {
	return r.WithContext(ContextWithUserID(r.Context(), userID))
}

func UserIDFromContext(ctx context.Context) int64 {
	return ctx.Value("user_id").(int64)
}
