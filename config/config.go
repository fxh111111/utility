package config

import "context"

var p provider

func GetString(ctx context.Context, key, dft string) string {
	if key == "" || p == nil {
		return dft
	}
	return p.getString(ctx, key, dft)
}
