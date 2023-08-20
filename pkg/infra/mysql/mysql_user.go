package mysql

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Insert(ctx context.Context, name string, token string) (string, error) {
	sql := `insert into users (name, token) values (:name, :token);`

	in := entity.User{
		Name:  name,
		Token: token,
	}

	_, err := ur.DB.NamedExec(sql, in)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ur *userRepository) SelectNameByUserID(ctx context.Context, userID int64) (string, error) {
	var user entity.User
	err := ur.DB.Get(&user, "select * from users where id = ?;", userID)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (ur *userRepository) SelectIDByToken(ctx context.Context, token string) (int64, error) {
	var user entity.User
	err := ur.DB.Get(&user, "select * from users where token = ?;", token)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (ur *userRepository) Update(ctx context.Context, name string, userID int64) error {
	sql := `update users set name = :name where id = :id;`

	in := entity.User{
		ID:   userID,
		Name: name,
	}

	_, err := ur.DB.NamedExec(sql, in)
	if err != nil {
		return err
	}

	return nil
}
