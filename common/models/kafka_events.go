package models

type UserMessageEvent struct {
	UserID  int      `json:"user_id"`
	Product []string `json:"product"`
}
