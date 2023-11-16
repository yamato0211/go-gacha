package redis

import (
	"context"
	"go-gacha-system/pkg/domain/repository"

	"github.com/go-redis/redis/v8"
)

const (
	redisRanking string = "RedisRanking"
)

type rankingRepository struct {
	client *redis.Client
}

func NewRankingRepository(c *redis.Client) repository.RankingRepository {
	return &rankingRepository{
		client: c,
	}
}

func (rr *rankingRepository) InsertAll(ctx context.Context, items []*redis.Z) error {
	if err := rr.client.ZAdd(ctx, redisRanking, items...).Err(); err != nil {
		return err
	}

	return nil
}

func (rr *rankingRepository) Select(ctx context.Context, limit int64) ([]redis.Z, error) {
	var start int64 = 0
	stop := start + limit - 1
	serializedMembersWithScores, err := rr.client.ZRevRangeWithScores(ctx, redisRanking, start, stop).Result()
	if err != nil {
		return nil, err
	}

	return serializedMembersWithScores, nil
}
