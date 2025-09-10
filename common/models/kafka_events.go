package models

type UserMessageEvent struct {
	UserID  int      `json:"user_id"`
	FavoriteProducts []string `json:"favorite_products"`
}

type ProductMessageEvent struct {
	ProductID int `json:"product_id"`
	ProductTags []string `json:"product_tags"`
}
