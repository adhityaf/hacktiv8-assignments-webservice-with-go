package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"foreignKey:OrderID;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
}
