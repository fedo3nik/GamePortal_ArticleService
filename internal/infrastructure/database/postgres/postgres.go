package postgres

import (
	"context"
	"log"

	"github.com/fedo3nik/GamePortal_ArticleService/internal/domain/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Insert(ctx context.Context, p *pgxpool.Pool, a *entities.Article) (int, error) {
	var id int

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return 0, err
	}

	defer func() {
		err := conn.Release
		if err != nil {
			log.Print("Release connection error")
		}
	}()

	row := conn.QueryRow(ctx, "INSERT INTO articles (userId, title, game, article_text, rating) VALUES($1, $2, $3, $4, $5) RETURNING id",
		&a.UserID, &a.Title, &a.Game, &a.Text, &a.Rating)

	err = row.Scan(&id)
	if err != nil {
		log.Printf("Scan error: %v", err)
	}

	return id, nil
}

func Select(ctx context.Context, p *pgxpool.Pool, id int) (*entities.Article, error) {
	var a entities.Article

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return nil, err
	}

	defer func() {
		err := conn.Release
		if err != nil {
			log.Print("Release connection error")
		}
	}()

	err = conn.QueryRow(ctx, "SELECT id, userId, title, game, article_text, rating FROM articles where id=$1", id).
		Scan(&a.ID, &a.UserID, &a.Title, &a.Game, &a.Text, &a.Rating)
	if err != nil {
		log.Printf("Select error: %v", err)
		return nil, err
	}

	return &a, nil
}
