package dtodb

type ArticleDTO struct {
	ID     int
	UserID string
	Title  string
	Game   string
	Text   string
}

type RatingDTO struct {
	ID           int
	ArticleID    int
	Rating       float64
	CountOfMarks int
}
