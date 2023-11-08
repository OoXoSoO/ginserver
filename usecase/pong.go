package usecase

import "context"

type PongService struct {
}

func (ps *PongService) Pong(ctx context.Context) string {
	return "pong"
}
