package models

type Item struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	ItemCode    string `gorm:"not null; unique" json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint
}
