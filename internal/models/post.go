package models

import "time"

type Post struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Content             string    `json:"content"`
	UserID              uint      `json:"userId"`
	CreatedAt           time.Time `json:"createdAt"`
	AreCommentsDisabled bool      `json:"areCommentsDisabled"`
}
