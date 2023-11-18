package config

import "context"

type KV interface {
	Get(ctx context.Context, key string) (string, error)
}

type kvProvider struct {
	client KV
}

func (p *kvProvider) getString(ctx context.Context, key string, defaultValue string) string {
	if p == nil || p.client == nil || key == "" {
		return defaultValue
	}
	s, err := p.client.Get(ctx, key)
	if err != nil {
		return defaultValue
	}
	return s
}

func SetKVProvider(kv KV) {
	p = &kvProvider{client: kv}
}
