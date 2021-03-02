package dtodb

type ArticleDTO struct {
	ID     int `sql:"int"`
	UserID string
	Title  string
	Game   string
	Rating float64
	Text   string
}
