package handlers

import (
	"github.com/n1kos03/User_Recommendations/internal/database"
)

type RecommendationService struct {
	DB *database.Database
}

func NewRecommendationService(db *database.Database) *RecommendationService {
	return &RecommendationService{DB: db}
}