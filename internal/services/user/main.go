package main

import (
	"User_Recommendations/internal/database"
	"User_Recommendations/internal/services/user/handlers"
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := database.DBConfig{
		User: "user",
		Pass: "password",
		Name: "user_db",
		Host: "localhost",
		Port: "5432",
		SSLMode: "disable",
	}

	DB, err := database.NewDatabase(cfg)
	if err != nil {
		slog.Error("Error connecting to database", "Error: ", err)
	}
	defer DB.Conn.Close()

	DB.CheckDBConnection()

	DB.RunMigrations()

	userHandlers := handlers.NewUserHandler(DB)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", userHandlers.GETUser)
	router.POST("/users", userHandlers.POSTUser)

	router.Run(":8080")
}