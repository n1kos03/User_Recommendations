package handlers

import (
	"github.com/n1kos03/User_Recommendations/common/database"
)

type ProductService struct {
	DB *database.Database
	// Producer *kafka.Producer
}

func NewProductService(db *database.Database /*, producer *kafka.Producer*/) *ProductService {
	return &ProductService{DB: db /*, Producer: producer*/}
}
