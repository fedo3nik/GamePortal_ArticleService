package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/pkg/errors"

	"github.com/fedo3nik/GamePortal_ArticleService/internal/application/service"
	dto "github.com/fedo3nik/GamePortal_ArticleService/internal/interface/controller/dtohttp"
	e "github.com/fedo3nik/GamePortal_ArticleService/internal/util/error"
)

type HTTPAddArticleHandler struct {
	articleService service.Article
}

type HTTPGetArticleHandler struct {
	articleService service.Article
}

func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, e.ErrDB) {
		_, hError := fmt.Fprintf(w, "Error caused: %v", err)
		if hError != nil {
			log.Printf("Fprint error: %v", hError)
		}

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, hError := fmt.Fprintf(w, "Internal server error: %v", err)
	if hError != nil {
		log.Printf("Fprint error: %v", hError)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func NewHTTPAddArticleHandler(articleService service.Article) *HTTPAddArticleHandler {
	return &HTTPAddArticleHandler{articleService: articleService}
}

func (hh HTTPAddArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dto.AddArticleRequest

	var resp dto.AddArticleResponse

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Body read error: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	article, err := hh.articleService.AddArticle(r.Context(), req.Title, req.Game, req.Text, req.Token)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.ID = article.ID
	resp.UserID = article.UserID
	resp.Game = article.Game
	resp.Title = article.Title

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		log.Printf("Encode error: %v", err)
		return
	}
}

func NewHTTPGetArticleHandler(articleService service.Article) *HTTPGetArticleHandler {
	return &HTTPGetArticleHandler{articleService: articleService}
}

func (hh HTTPGetArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp dto.GetArticleResponse

	url := r.URL.Path
	idString := path.Base(url)

	article, err := hh.articleService.GetArticle(r.Context(), idString)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.ID = article.ID
	resp.Title = article.Title
	resp.Game = article.Game
	resp.UserID = article.UserID
	resp.Text = article.Text
	resp.Ratting = article.Rating

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		handleError(w, err)
		return
	}
}
