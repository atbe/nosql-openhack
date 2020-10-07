package category

type Category struct {
	ID          string `json:"id" gorm:"column:CategoryId"`
	Name        string `json:"name" gorm:"column:CategoryName"`
	Description string `json:"description" gorm:"column:Description"`
}

func (Category) TableName() string {
	return "Category"
}
