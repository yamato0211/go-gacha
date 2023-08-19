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

func (ur *userRepository) GetByName(ctx context.Context, token string) (string, error) {
	var user entity.User
	err := ur.DB.Get(&user, "select * from users where token = ?;", token)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (ur *userRepository) Update(ctx context.Context, name string, token string) error {
	sql := `update users set name = :name where token = :token;`

	in := entity.User{
		Token: token,
		Name:  name,
	}

	_, err := ur.DB.NamedExec(sql, in)
	if err != nil {
		return err
	}

	return nil
}
