package models

type UserMessageEvent struct {
	UserID  int      `json:"user_id"`
	FavoriteProducts []string `json:"product"`
}
