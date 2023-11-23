package pkg

import (
	"context"
)

type PongService interface {
	Pong(ctx context.Context) (string, error)
}
