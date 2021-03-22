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
	SetMark(ctx context.Context, title string, mark float64) (*entities.Article, error)
}

type ArticleService struct {
	Pool       *pgxpool.Pool
	AccessKey  string
	RefreshKey string
}

func (a ArticleService) AddArticle(ctx context.Context, title, game, text, token string) (*entities.Article, error) {
	userID, err := service.ValidateAccessToken(token, a.AccessKey)
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

func (a ArticleService) SetMark(ctx context.Context, title string, mark float64) (*entities.Article, error) {
	articleID, err := postgres.InsertRating(ctx, a.Pool, mark, title)
	if err != nil {
		return nil, err
	}

	article, err := postgres.Select(ctx, a.Pool, articleID)
	if err != nil {
		return nil, errors.Wrap(e.ErrDB, "select")
	}

	return article, nil
}

func NewArticleService(pool *pgxpool.Pool, accessKey, refreshKey string) *ArticleService {
	return &ArticleService{Pool: pool, AccessKey: accessKey, RefreshKey: refreshKey}
}
