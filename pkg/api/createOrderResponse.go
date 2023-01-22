package api

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrderResponse struct {
	Id          uuid.UUID     `json:"id"`
	UserId      uuid.UUID     `json:"userId"`
	Items       []RequestItem `json:"items"`
	CreatedDate time.Time     `json:"createdDate"`
}
