package repository

import (
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, name string, token string) (string, error)
	SelectNameByUserID(ctx context.Context, userID int64) (string, error)
	SelectIDByToken(ctx context.Context, token string) (int64, error)
	Update(ctx context.Context, name string, userID int64) error
}
