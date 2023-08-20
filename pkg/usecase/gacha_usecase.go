package usecase

import "context"

type IGachaUsecase interface {
	DrawGacha(ctx context.Context, userID int64, characterID int64)
}
