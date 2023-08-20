package mysql

import (
	"context"
	"go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type characterRepository struct {
	DB *sqlx.DB
}

func NewCharacterRepository(db *sqlx.DB) repository.CharacterRepository {
	return &characterRepository{
		DB: db,
	}
}

func (cr *characterRepository) SelectAllByID(ctx context.Context, userID int64) ([]entity.Character, error) {
	var characters []entity.Character
	query := `
		select
		    c.id, 
		    c.name,
		    c.description,
		    c.hp,
		    c.cost,
		    c.power,
		    c.speed,
		    c.rarity,
		    cu.count
		from
		    characters_users cu
		inner join
		    characters c ON cu.character_id = c.id
		where cu.user_id = ?;
	`
	if err := cr.DB.Select(&characters, query, userID); err != nil {
		return nil, err
	}

	return characters, nil
}
