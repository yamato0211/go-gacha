package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RankingRepository interface {
	Select(ctx context.Context, limit int64) ([]redis.Z, error)
	InsertAll(ctx context.Context, items []*redis.Z) error
}
