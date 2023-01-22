package api

import "github.com/google/uuid"

type RequestItem struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
