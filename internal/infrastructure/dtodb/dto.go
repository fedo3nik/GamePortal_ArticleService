package dtodb

type ArticleDTO struct {
	ID     int
	UserID string
	Title  string
	Game   string
	Text   string
}

type RattingDTO struct {
	ID           int
	ArticleID    int
	Ratting      float64
	CountOfMarks int
}
