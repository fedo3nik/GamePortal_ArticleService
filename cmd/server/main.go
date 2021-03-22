package main

import (
	"context"
	"log"
	"net/http"

	grpcInfra "github.com/fedo3nik/GamePortal_ArticleService/internal/infrastructure/grpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"

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

	grpcConn, err := grpc.Dial(c.GrpcPort, grpc.WithInsecure())
	if err != nil {
		log.Panicf("Grpc connection error: %v", err)
	}

	grpcClient := grpcInfra.NewSenderClient(grpcConn)

	emp := grpcInfra.Empty{}

	grpcResp, err := grpcClient.Send(context.Background(), &emp)
	if err != nil {
		log.Panicf("Grpc received error: %v", err)
	}

	handler := mux.NewRouter()

	articleService := service.NewArticleService(pool, grpcResp.AccessPublicKey, grpcResp.RefreshPublicKey)
	addArticleHandler := controller.NewHTTPAddArticleHandler(articleService)
	getArticleHandler := controller.NewHTTPGetArticleHandler(articleService)
	setMarkHandler := controller.NewHTTPSetMarkToArticleHandler(articleService)

	handler.Handle("/new-article", addArticleHandler).Methods("POST")
	handler.Handle("/article/{articleId}", getArticleHandler).Methods("GET")
	handler.Handle("/article/set-mark", setMarkHandler).Methods("POST")

	go func() {
		err = http.ListenAndServe(c.Host+c.Port, handler)
		if err != nil {
			log.Panicf("Listen & Serve serror: %v", err)
		}
	}()

	select {}
}
