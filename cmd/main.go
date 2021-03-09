package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/fedo3nik/GamePortal_ArticleService/internal/application/service"
	"github.com/fedo3nik/GamePortal_ArticleService/internal/config"
	"github.com/fedo3nik/GamePortal_ArticleService/internal/interface/controller"
	"github.com/gorilla/mux"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Panicf("Config create error: %v", err)
	}

	pool, err := pgxpool.Connect(context.Background(), c.PostgresURL)
	if err != nil {
		log.Panicf("Connect error: %v", err)
	}

	defer pool.Close()

	handler := mux.NewRouter()

	articleService := service.NewArticleService(pool)
	addArticleHandler := controller.NewHTTPAddArticleHandler(articleService)
	getArticleHandler := controller.NewHTTPGetArticleHandler(articleService)

	handler.Handle("/new-article", addArticleHandler).Methods("POST")
	handler.Handle("/article/{articleId}", getArticleHandler).Methods("GET")

	err = http.ListenAndServe(c.Host+c.Port, handler)
	if err != nil {
		log.Panicf("Listen & Serve serror: %v", err)
	}
}
