package models

import "time"

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	FavoriteProduct string `json:"favorite_product"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}