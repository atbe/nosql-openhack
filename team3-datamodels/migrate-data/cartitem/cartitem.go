package cartitem

type CartItem struct {
	CartItemId string `json:"cart_item_id" gorm:"column:CartItemId;`
	ItemId     string `json:"ItemId`
	Quantity   uint   `json:"Quantity`
}

func (CartItem) TableName() string {
	return "CartItem"
}
