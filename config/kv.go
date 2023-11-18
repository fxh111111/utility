package config

import "context"

type KV interface {
	Get(ctx context.Context, key string) (string, error)
}

type KVProvider struct {
	client KV
}

func (p *KVProvider) GetString(ctx context.Context, key string, defaultValue string) string {
	s, err := p.client.Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	return s
}
