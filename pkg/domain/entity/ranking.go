package entity

type Ranking struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
	Rank  int    `json:"rank"`
}
