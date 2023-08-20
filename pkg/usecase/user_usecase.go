package usecase

import (
	"context"
	"go-gacha-system/pkg/domain/repository"

	"github.com/google/uuid"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, name string) (string, error)
	UpdateName(ctx context.Context, name string, userID int64) error
	GetUser(ctx context.Context, userID int64) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uu *userUsecase) CreateUser(ctx context.Context, name string) (string, error) {
	uuid := uuid.New().String()
	token, err := uu.repo.Insert(ctx, name, uuid)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uu *userUsecase) GetUser(ctx context.Context, userID int64) (string, error) {
	name, err := uu.repo.SelectNameByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (uu *userUsecase) UpdateName(ctx context.Context, name string, userID int64) error {
	err := uu.repo.Update(ctx, name, userID)
	if err != nil {
		return err
	}

	return nil
}
