package util

import (
	"context"
	"net/http"
)

var (
	userIdKey = "user_id"
)

func ContextWithUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIdKey, userID)
}

func RequestWithUserID(r *http.Request, userID uint) *http.Request {
	return r.WithContext(ContextWithUserID(r.Context(), userID))
}

func UserIDFromContext(ctx context.Context) uint {
	return ctx.Value("user_id").(uint)
}
