package storage

import "context"

type Storage interface {
	Save(ctx context.Context, original string) (string, error)
	Get(ctx context.Context, code string) (string, error)
	Close() error
}
