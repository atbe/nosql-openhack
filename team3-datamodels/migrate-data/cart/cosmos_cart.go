package cart

import "openhack/cartitem"

type CosmosCart struct {
	UserId string                    `json:"UserId"`
	Id     string                    `json:"Id"`
	Items  []cartitem.CosmosCartItem `json:"Items"`
}
