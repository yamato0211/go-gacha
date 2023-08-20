package usecase

import (
	"context"
	"database/sql"
	"errors"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"
	"math/rand"
	"time"
)

type IGachaUsecase interface {
	DrawGacha(ctx context.Context, userID int64, count int64) ([]entity.Character, error)
}

type gachaUsecase struct {
	gachaRepo     repository.GachaRepository
	characterRepo repository.CharacterRepository
	userRepo      repository.UserRepository
}

func NewGachaUsecase(gr repository.GachaRepository, cr repository.CharacterRepository, ur repository.UserRepository) IGachaUsecase {
	return &gachaUsecase{
		gachaRepo:     gr,
		characterRepo: cr,
		userRepo:      ur,
	}
}

func (gu *gachaUsecase) DrawGacha(ctx context.Context, userID int64, count int64) ([]entity.Character, error) {
	var totalProbability = 1000
	characters := make([]entity.Character, 0, count)
	contents, err := gu.gachaRepo.SelectAll(ctx)
	if err != nil {
		return nil, errors.New("failed SelectAll")
	}

	for i := 0; i < int(count); i++ {
		//乱数の作成
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(totalProbability)
		var chosenID int64
		for _, content := range contents {
			num -= int(content.Probability)
			if num < 0 {
				chosenID = content.CharacterID
				break
			}
		}

		err = gu.gachaRepo.SelectCountByUserID(ctx, userID, chosenID)
		switch err {
		case nil:
			err = gu.gachaRepo.UpdateByUserID(ctx, userID, chosenID)
			if err != nil {
				return nil, errors.New("failed UpdateByUserID")
			}
		case sql.ErrNoRows:
			err = gu.gachaRepo.InsertByUserID(ctx, userID, chosenID)
			if err != nil {
				return nil, errors.New("failed InsertByUserID")
			}
		default:
			return nil, errors.New("default error")
		}

		chara, err := gu.characterRepo.SelectByUserID(ctx, userID, chosenID)
		if err != nil {
			return nil, errors.New("failed SelectByUserID")
		}
		characters = append(characters, chara)
	}

	return characters, nil
}
