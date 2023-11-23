package service

import (
	"context"
	"ginserver/pkg"
)

type Pong struct {
}

// ensure pkg implementation
var _ pkg.PongService = (*Pong)(nil)

func (ps *Pong) Pong(ctx context.Context) (string, error) {
	return "pong", nil
}
