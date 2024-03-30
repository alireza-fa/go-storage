package models

type User struct {
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
	CreatedAt    string `json:"created_at"`
}
