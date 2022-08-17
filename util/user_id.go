package util

import (
	"context"
	"net/http"
)

var (
	userIdKey = "user_id"
	roleIdKey = "role_id"
)

func ContextWithUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIdKey, userID)
}

func RequestWithUserID(r *http.Request, userID uint) *http.Request {
	return r.WithContext(ContextWithUserID(r.Context(), userID))
}

func UserIDFromContext(ctx context.Context) uint {
	return ctx.Value(userIdKey).(uint)
}


func UserIDFromContext2(ctx context.Context) interface{} {
	return ctx.Value(userIdKey)
}

func ContextWithRoleID(ctx context.Context, roleID uint) context.Context {
	return context.WithValue(ctx, roleIdKey, roleID)
}

func RequestWithRoleID(r *http.Request, roleID uint) *http.Request {
	return r.WithContext(ContextWithRoleID(r.Context(), roleID))
}

func RoleIDFromContext(ctx context.Context) uint {
	return ctx.Value(roleIdKey).(uint)
}
