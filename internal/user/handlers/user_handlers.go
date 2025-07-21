package handlers

import (
	"github.com/n1kos03/User_Recommendations/internal/database"
	"github.com/n1kos03/User_Recommendations/internal/kafka"
)

type UserHandler struct {
	DB       *database.Database
	Producer *kafka.Producer
}

func NewUserHandler(db *database.Database, producer *kafka.Producer) *UserHandler {
	return &UserHandler{DB: db, Producer: producer}
}
