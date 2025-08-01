package models

import "time"

type Recommendation struct{
	ID int `json:"id"`
	UserID int `json:"user_id"`
	ProductID int `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}