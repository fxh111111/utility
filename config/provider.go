package config

import "context"

type provider interface {
	getString(ctx context.Context, key string, dft string) string
}
