package session

import (
	"context"

	"github.com/deerling/resourcepack/internal/models"
	"github.com/sirupsen/logrus"
)

type contextValueKey int

const (
	contextKeyUser contextValueKey = iota
	contextKeyLogger
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

func Logger(ctx context.Context) *logrus.Entry {
	v, _ := ctx.Value(contextKeyLogger).(*logrus.Entry)
	return v
}

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, contextKeyLogger, logger)
}
