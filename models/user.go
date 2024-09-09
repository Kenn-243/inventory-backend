package models

type User struct {
	UserID int `json:"userId"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}