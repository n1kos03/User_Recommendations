package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/n1kos03/User_Recommendations/common/database"
	"github.com/n1kos03/User_Recommendations/common/kafka"
	"github.com/n1kos03/User_Recommendations/services/product/handlers"
)

func main() {
	cfg := database.DBConfig{
		User:    "admin",
		Pass:    "password",
		Name:    "product_db",
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

	DB.RunMigrations("file://common/database/migrations/product-service", cfg.Name)

	producer, err := kafka.NewProducer([]string{"kafka:9092"})
	if err != nil {
		slog.Error("Error while creating new instance for producer", "Error: ", err)
	}
	defer producer.Close()
	
	productService := handlers.NewProductService(DB, producer)

	router := gin.Default()

	router.POST("/products", productService.POSTProduct)
	router.GET("/products/:id", productService.GETProduct)
	router.GET("/products", productService.GETProductByTag)
	router.PUT("/products/:id", productService.PUTProduct)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Error running server", "Error: ", err)
	}
}
