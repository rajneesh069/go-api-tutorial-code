package handlers

import (
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajneesh069/go-api-tutorial-code/internal/database"
	"github.com/rajneesh069/go-api-tutorial-code/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func CreateUserHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := CreateUserInput{}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if utf8.RuneCountInString(input.Password) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is too short, it should be atleast 8 characters"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
			return
		}

		createdUser, err := repository.CreateUser(c, pool, input.Email, string(hashedPassword))

		if err != nil {
			if database.IsUniqueConstraintViolation(err) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email registered already"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, createdUser)
	}
}
