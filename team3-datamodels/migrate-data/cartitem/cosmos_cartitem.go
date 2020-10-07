package cartitem

import "openhack/item"

type CosmosCartItem struct {
	Quantity  uint
	UnitPrice float32
	Item      item.CosmosItem `json:"Item"`
}
