package logger

import "context"

type Logger interface {
	Error(ctx context.Context, a ...any)
}
