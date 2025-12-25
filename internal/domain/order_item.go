package domain

import "time"

type OrderItem struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"orderId"`
	ProductID int64     `json:"productId"`
	Quantity  int64     `json:"quantity"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
