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

	defer conn.Release()

	row := conn.QueryRow(ctx, "INSERT INTO articles (userId, title, game, article_text) VALUES($1, $2, $3, $4) RETURNING id",
		&a.UserID, &a.Title, &a.Game, &a.Text)

	err = row.Scan(&id)
	if err != nil {
		log.Printf("Scan error: %v", err)
		return 0, err
	}

	return id, nil
}

func Select(ctx context.Context, p *pgxpool.Pool, id int) (*entities.Article, error) {
	var a entities.Article

	var articleRatting float64

	var countOfMarks int

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return nil, err
	}

	defer conn.Release()

	err = conn.QueryRow(ctx, "SELECT id, userId, title, game, article_text FROM articles WHERE id=$1", id).
		Scan(&a.ID, &a.UserID, &a.Title, &a.Game, &a.Text)
	if err != nil {
		log.Printf("Select error: %v", err)
		return nil, err
	}

	row := conn.QueryRow(ctx, "SELECT COUNT(*) FROM rating r INNER JOIN articles a ON r.articleID=a.ID WHERE a.ID=$1", id)

	err = row.Scan(&countOfMarks)
	if err != nil {
		log.Printf("Select error: %v", err)
		return nil, err
	}

	if countOfMarks == 0 {
		a.Rating = 0
		return &a, nil
	}

	row = conn.QueryRow(ctx, "SELECT SUM(rating) as totalRatting FROM rating r INNER JOIN articles a ON r.articleID=a.ID WHERE a.ID=$1", id)

	err = row.Scan(&articleRatting)
	if err != nil {
		log.Printf("Select error: %v", err)
		return nil, err
	}

	a.Rating = articleRatting / float64(countOfMarks)

	return &a, nil
}

func InsertRating(ctx context.Context, p *pgxpool.Pool, mark float64, title string) (int, error) {
	var articleID int

	var id int

	conn, err := p.Acquire(ctx)
	if err != nil {
		log.Printf("Create connection from pool error: %v", err)
		return 0, err
	}

	defer conn.Release()

	err = conn.QueryRow(ctx, "SELECT id FROM articles WHERE title=$1", title).
		Scan(&articleID)
	if err != nil {
		log.Printf("Select error: %v", err)
		return 0, err
	}

	err = conn.QueryRow(ctx, "INSERT INTO Rating (articleID, rating) VALUES($1, $2) RETURNING id", articleID, mark).
		Scan(&id)
	if err != nil {
		log.Printf("Select error: %v", err)
		return 0, err
	}

	return articleID, nil
}
