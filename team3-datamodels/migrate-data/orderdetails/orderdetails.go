package orderdetails

type OrderDetails struct {
	ProductID string  `json:"product_id" gorm:"column:ProductId"`
	UnitPrice float32 `json:"unit_price" gorm:"column:UnitPrice"`
}

func (OrderDetails) TableName() string {
	return "OrderDetails"
}
