package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajneesh069/go-api-tutorial-code/internal/models"
)

func CreateTodo(ctx context.Context, pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO todos (title, completed)
		VALUES ($1, $2)
		RETURNING id, title, completed, created_at, updated_at
	`

	var todo models.Todo

	err := pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &todo, nil
}
