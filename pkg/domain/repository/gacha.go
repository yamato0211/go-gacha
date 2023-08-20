package repository

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
)

type GachaRepository interface {
	SelectAll(ctx context.Context) ([]entity.GachaContent, error)
	SelectCountByUserID(ctx context.Context, userID int64, characterID int64) error
	InsertByUserID(ctx context.Context, userID int64, characterID int64) error
	UpdateByUserID(ctx context.Context, userID int64, characterID int64) error
}
