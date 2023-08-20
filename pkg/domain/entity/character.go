package entity

type Character struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Hp          int64  `json:"hp" db:"hp"`
	Cost        int64  `json:"cost" db:"cost"`
	Power       int64  `json:"power" db:"power"`
	Speed       int64  `json:"speed" db:"speed"`
	Rarity      int64  `json:"rarity" db:"rarity"`
	Count       int64  `json:"count" db:"count"`
}
