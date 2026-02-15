package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rajneesh069/go-api-tutorial-code/internal/config"
	"github.com/rajneesh069/go-api-tutorial-code/internal/database"
	"github.com/rajneesh069/go-api-tutorial-code/internal/handlers"
)

func main() {
	var cfg *config.Config
	var err error

	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer pool.Close()
	var router *gin.Engine = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   200,
			"message":  "Todo API is running",
			"database": "connected",
		})
	})

	router.POST("/todos", handlers.CreateTodoHandler(pool))

	router.Run(":" + cfg.Port)

}
