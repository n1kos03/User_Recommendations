package handlers

import "User_Recommendations/internal/database"

type UserHandler struct {
	DB *database.Database
}

func NewUserHandler(db *database.Database) *UserHandler {
	return &UserHandler{DB: db}
}