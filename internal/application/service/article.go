package service

import (
	"context"
	"log"
	"strconv"

	service "github.com/fedo3nik/GamePortal_ArticleService/internal/application"

	"github.com/pkg/errors"

	"github.com/fedo3nik/GamePortal_ArticleService/internal/domain/entities"
	"github.com/fedo3nik/GamePortal_ArticleService/internal/infrastructure/database/postgres"
	e "github.com/fedo3nik/GamePortal_ArticleService/internal/util/error"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Article interface {
	AddArticle(ctx context.Context, title string, game string, text string, token string) (*entities.Article, error)
	GetArticle(ctx context.Context, ID string) (*entities.Article, error)
}

type ArticleService struct {
	Pool *pgxpool.Pool
}

func (a ArticleService) AddArticle(ctx context.Context, title, game, text, token string) (*entities.Article, error) {
	userID, err := service.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}

	var article entities.Article
	article.Title = title
	article.Game = game
	article.Text = text
	article.Rating = 0
	article.UserID = userID

	id, err := postgres.Insert(ctx, a.Pool, &article)
	if err != nil {
		log.Printf("DB: %v", err)
		return nil, errors.Wrap(e.ErrDB, "insert")
	}

	article.ID = id

	return &article, nil
}

func (a ArticleService) GetArticle(ctx context.Context, strID string) (*entities.Article, error) {
	id, err := strconv.Atoi(strID)
	if err != nil {
		log.Printf("Convert error: %v", err)
		return nil, err
	}

	art, err := postgres.Select(ctx, a.Pool, id)
	if err != nil {
		return nil, errors.Wrap(e.ErrDB, "select")
	}

	return art, nil
}

func NewArticleService(pool *pgxpool.Pool) *ArticleService {
	return &ArticleService{Pool: pool}
}
