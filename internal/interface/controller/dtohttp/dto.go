package controller

type AddArticleRequest struct {
	Title string `json:"title"`
	Game  string `json:"game"`
	Text  string `json:"text"`
	Token string `json:"token"`
}

type AddArticleResponse struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Game   string `json:"game"`
}

type GetArticleResponse struct {
	ID     int     `json:"id"`
	UserID string  `json:"user_id"`
	Title  string  `json:"title"`
	Game   string  `json:"game"`
	Text   string  `json:"text"`
	Rating float64 `json:"rating"`
}
