package repository

import (
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, name string, token string) (string, error)
	SelectNameByToken(ctx context.Context, token string) (string, error)
	SelectIDByToken(ctx context.Context, token string) (int64, error)
	Update(ctx context.Context, name string, token string) error
}
