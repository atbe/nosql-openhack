package cart

type Cart struct {
	UserId string `json:"user_id" gorm:"column:UserId"`
	ID     string `json:"id" gorm:"column:Id"`
}

func (Cart) TableName() string {
	return "Cart"
}
