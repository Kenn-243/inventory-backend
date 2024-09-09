package models

type Item struct {
	ItemID int `json:"itemId"`
	ItemName string `json:"itemName"`
	UserID int `json:"userId"`
}