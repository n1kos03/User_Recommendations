package main

import (
	"net/http"

	"github.com/n1kos03/User_Recommendations/internal/database"
	"github.com/n1kos03/User_Recommendations/internal/kafka"
	"github.com/n1kos03/User_Recommendations/internal/user/handlers"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := database.DBConfig{
		User:    "user",
		Pass:    "password",
		Name:    "user_db",
		Host:    "postgres",
		Port:    "5432",
		SSLMode: "disable",
	}

	DB, err := database.NewDatabase(cfg)
	if err != nil {
		slog.Error("Error connecting to database", "Error: ", err)
	}

	defer func() {
		if err = DB.Conn.Close(); err != nil {
			slog.Error("Error closing database connection", "Error: ", err)
		}
	}()

	DB.CheckDBConnection()

	DB.RunMigrations()

	produser, err := kafka.NewProducer([]string{"kafka:9092"})
	if err != nil {
		slog.Error("Error creating producer", "Error: ", err)
	}
	defer produser.Close()

	userService := handlers.NewUserService(DB, produser)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", userService.GETUser)
	router.POST("/users", userService.POSTUser)
	router.PUT("/users/:id", userService.PUTUser)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Error running server", "Error: ", err)
	}
}
