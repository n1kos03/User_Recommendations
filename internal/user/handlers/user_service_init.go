package handlers

import (
	"github.com/n1kos03/User_Recommendations/internal/database"
	"github.com/n1kos03/User_Recommendations/internal/kafka"
)

type UserService struct {
	DB       *database.Database
	Producer *kafka.Producer
}

func NewUserService(db *database.Database, producer *kafka.Producer) *UserService {
	return &UserService{DB: db, Producer: producer}
}
