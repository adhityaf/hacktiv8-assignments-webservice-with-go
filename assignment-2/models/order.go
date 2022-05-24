package models

import "time"

type Order struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CustomerName string    `gorm:"type:varchar(50)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item
}
