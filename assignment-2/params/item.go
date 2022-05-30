package params

type CreateItem struct {
	Description string `json:"description"`
	ItemCode    uint   `json:"item_code"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
