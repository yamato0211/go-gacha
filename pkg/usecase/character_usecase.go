package usecase

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"
)

type ICharacterUsecase interface {
	GetCharacters(ctx context.Context, userID int64) ([]entity.Character, error)
}

type characterUsecase struct {
	charaRepo repository.CharacterRepository
	userRepo  repository.UserRepository
}

func NewCharacterUsecase(cr repository.CharacterRepository, ur repository.UserRepository) ICharacterUsecase {
	return &characterUsecase{
		charaRepo: cr,
		userRepo:  ur,
	}
}

func (cu *characterUsecase) GetCharacters(ctx context.Context, userID int64) ([]entity.Character, error) {
	characters, err := cu.charaRepo.SelectAllByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return characters, nil
}
