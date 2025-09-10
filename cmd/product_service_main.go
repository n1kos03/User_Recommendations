package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/n1kos03/User_Recommendations/common/database"
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

	productService := handlers.NewProductService(DB)

	router := gin.Default()

	router.POST("/products", productService.POSTProduct)
	router.GET("/products/:id", productService.GETProduct)
	router.GET("/products", productService.GETProductByTag)
	router.PUT("/products/:id", productService.PUTProduct)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Error running server", "Error: ", err)
	}
}
