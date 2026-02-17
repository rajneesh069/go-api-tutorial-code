package database

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueConstraintViolation(err error) bool {
	pgErr := &pgconn.PgError{}
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
