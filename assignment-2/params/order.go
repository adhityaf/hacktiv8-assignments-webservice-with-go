package params

import "time"

type CreateOrder struct {
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}

type UpdateOrder struct {
	CustomerName string `json:"customer_name"`
}
