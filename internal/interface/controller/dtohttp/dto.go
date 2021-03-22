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

type SetMarkRequest struct {
	Title string  `json:"title"`
	Mark  float64 `json:"mark"`
}

type SetMarkResponse struct {
	ID     int     `json:"articleID"`
	Rating float64 `json:"rating"`
}
