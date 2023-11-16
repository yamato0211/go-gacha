package repository

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
)

type CharacterRepository interface {
	SelectAll(ctx context.Context) ([]*entity.Character, error)
	SelectAllOrderOffset(ctx context.Context, limit int64) ([]*entity.Character, error)
	SelectAllByID(ctx context.Context, userID int64) ([]entity.Character, error)
	SelectByUserID(ctx context.Context, userID int64, characterID int64) (entity.Character, error)
}
