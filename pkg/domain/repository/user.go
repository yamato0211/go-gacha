package repository

import (
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, name string, token string) (string, error)
	GetByName(ctx context.Context, token string) (string, error)
	Update(ctx context.Context, name string, token string) error
}
