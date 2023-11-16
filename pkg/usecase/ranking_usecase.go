package usecase

import (
	"context"
	"encoding/json"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"
	"log"

	"github.com/go-redis/redis/v8"
)

type IRankingUsecase interface {
	GetRedisRanking(ctx context.Context, limit int64) ([]*entity.Ranking, error)
	GetMysqlRanking(ctx context.Context, limit int64) ([]*entity.Ranking, error)
	InsertRedisRanking(ctx context.Context) error
}

type rankingUsecase struct {
	charaRepo   repository.CharacterRepository
	rankingRepo repository.RankingRepository
}

func NewRankingUsecase(cr repository.CharacterRepository, rr repository.RankingRepository) IRankingUsecase {
	return &rankingUsecase{
		charaRepo:   cr,
		rankingRepo: rr,
	}
}

func (ru *rankingUsecase) GetRedisRanking(ctx context.Context, limit int64) ([]*entity.Ranking, error) {
	serializedMembers, err := ru.rankingRepo.Select(ctx, limit)
	if err != nil {
		return nil, err
	}
	res := make([]*entity.Ranking, 0, limit)
	member := &entity.Character{}
	for i, serializedMemberWithScore := range serializedMembers {
		serializedMember := serializedMemberWithScore.Member
		score := serializedMemberWithScore.Score
		if err := json.Unmarshal([]byte(serializedMember.(string)), member); err != nil {
			return nil, err
		}
		r := &entity.Ranking{
			ID:    member.ID,
			Name:  member.Name,
			Score: int64(score),
			Rank:  i + 1,
		}
		res = append(res, r)
	}

	return res, nil
}

func (ru *rankingUsecase) GetMysqlRanking(ctx context.Context, limit int64) ([]*entity.Ranking, error) {
	charactes, err := ru.charaRepo.SelectAllOrderOffset(ctx, limit)
	if err != nil {
		return nil, err
	}

	rankings := make([]*entity.Ranking, 0, limit)
	for i, character := range charactes {
		ranking := &entity.Ranking{
			ID:    character.ID,
			Name:  character.Name,
			Score: character.Power,
			Rank:  i + 1,
		}
		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

func (ru *rankingUsecase) InsertRedisRanking(ctx context.Context) error {
	characters, err := ru.charaRepo.SelectAll(ctx)
	if err != nil {
		log.Print("failt select all")
		return err
	}
	items := make([]*redis.Z, 0, len(characters))
	for _, member := range characters {
		serializeMember, err := json.Marshal(member)
		if err != nil {
			return err
		}

		item := &redis.Z{
			Score:  float64(member.Power),
			Member: serializeMember,
		}
		items = append(items, item)
	}
	if err := ru.rankingRepo.InsertAll(ctx, items); err != nil {
		log.Print("failt insert all")
		return err
	}

	return nil
}
