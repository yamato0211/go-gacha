package entity

type GachaContent struct {
	CharacterID int64 `json:"character_id" db:"character_id"`
	Probability int64 `json:"probability" db:"probability"`
}

type CharacterUser struct {
	UserID      int `db:"user_id"`
	CharacterID int `db:"character_id"`
	Count       int `db:"count"`
}
