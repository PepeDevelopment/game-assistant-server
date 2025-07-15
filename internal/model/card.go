package model

type Card struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
