package config

import "context"

type Provider interface {
	GetString(ctx context.Context, key string, dft string) string
}
