package session

import (
	"context"

	"github.com/deerling/resources.app/internal/models"
)

type contextValueKey int

const (
	contextKeyUser contextValueKey = iota
	contextKeyDatabase
)

// User reads User from context
func User(ctx context.Context) *models.User {
	v, _ := ctx.Value(contextKeyUser).(*models.User)
	return v
}

// WithUser puts User into context
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, contextKeyUser, user)
}
