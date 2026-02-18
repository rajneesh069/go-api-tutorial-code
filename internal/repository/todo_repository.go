package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajneesh069/go-api-tutorial-code/internal/models"
)

func CreateTodo(ctx context.Context, pool *pgxpool.Pool, title string, completed bool, userId string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO todos (title, completed, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, completed, created_at, updated_at, user_id
	`

	var todo models.Todo

	err := pool.QueryRow(ctx, query, title, completed, userId).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserId,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &todo, nil
}

func GetAllTodos(ctx context.Context, pool *pgxpool.Pool, userId string) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT id, title, completed, created_at, updated_at, user_id 
		FROM todos
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := pool.Query(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []models.Todo{}

	for rows.Next() {
		todo := models.Todo{}

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.UserId,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoByID(ctx context.Context, pool *pgxpool.Pool, id uint, userId string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	todo := models.Todo{}
	query := `
		SELECT id, title, completed, created_at, updated_at, user_id
		FROM todos
		WHERE id = $1 AND user_id = $2
	`

	err := pool.QueryRow(ctx, query, id, userId).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserId,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(ctx context.Context, pool *pgxpool.Pool, id uint, title string, completed bool, userId string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	todo := models.Todo{}
	query := `
		UPDATE todos 
		SET title = $1, completed = $2
		WHERE id = $3 AND user_id = $4
		RETURNING id, title, completed, created_at, updated_at, user_id
	`
	err := pool.QueryRow(ctx, query, title, completed, id, userId).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserId,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(ctx context.Context, pool *pgxpool.Pool, id uint, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		DELETE FROM todos
		WHERE id = $1 AND user_id = $2
	`
	commandTag, err := pool.Exec(ctx, query, id, userId)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("TODO_NOT_FOUND")
	}

	return nil
}
