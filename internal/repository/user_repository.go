package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajneesh069/go-api-tutorial-code/internal/models"
)

func CreateUser(ctx context.Context, pool *pgxpool.Pool, email string, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := models.User{}

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, created_at, updated_at
	`
	err := pool.QueryRow(ctx, query, email, password).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(ctx context.Context, pool *pgxpool.Pool, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := models.User{}

	query := `
		SELECT id, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := models.User{}

	query := `
		SELECT id, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
