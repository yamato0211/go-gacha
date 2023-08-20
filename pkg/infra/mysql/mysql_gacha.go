package mysql

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type gachaRepository struct {
	DB *sqlx.DB
}

func NewGachaRepository(db *sqlx.DB) repository.GachaRepository {
	return &gachaRepository{
		DB: db,
	}
}

func (gr *gachaRepository) SelectAll(ctx context.Context) ([]entity.GachaContent, error) {
	var gachaContents []entity.GachaContent
	query := "select * from gacha_contents;"
	if err := gr.DB.Select(&gachaContents, query); err != nil {
		return nil, err
	}

	return gachaContents, nil
}

func (gr *gachaRepository) SelectCountByUserID(ctx context.Context, userID int64, characterID int64) error {
	var countResult entity.CharacterUser
	query := `SELECT count FROM characters_users WHERE user_id = ? AND character_id = ?;`
	err := gr.DB.Get(&countResult, query, userID, characterID)
	return err
}

func (gr *gachaRepository) InsertByUserID(ctx context.Context, userID int64, characterID int64) error {
	insertQuery := `INSERT INTO characters_users (user_id, character_id, count) VALUES (?, ?, 1);`
	_, err := gr.DB.Exec(insertQuery, userID, characterID)
	return err
}

func (gr *gachaRepository) UpdateByUserID(ctx context.Context, userID int64, characterID int64) error {
	updateQuery := `UPDATE characters_users SET count = count + 1 WHERE user_id = ? AND character_id = ?;`
	_, err := gr.DB.Exec(updateQuery, userID, characterID)
	return err
}
