package models

type Item struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	ItemCode    uint `gorm:"not null;" json:"item_code"`
	Description string `gorm:"not null; type:varchar(255)" json:"description"`
	Quantity    uint   `gorm:"not null; type:int" json:"quantity"`
	OrderID     uint   `gorm:"type:int" json:"order_id"`
}
