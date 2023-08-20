package repository

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
)

type CharacterRepository interface {
	SelectAllByID(ctx context.Context, userID int64) ([]entity.Character, error)
	SelectByUserID(ctx context.Context, userID int64, characterID int64) (entity.Character, error)
}
