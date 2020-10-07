package item

import (
	"time"

	"openhack/category"
)

type Item struct {
	ID          string            `json:"id" gorm:"column:ItemId"`
	Name        string            `json:"name" gorm:"column:ProductName"`
	ReleaseDate time.Time         `json:"name" gorm:"column:ReleaseDate"`
	UnitPrice   float32           `json:"unit_price" gorm:"column:UnitPrice"`
	CategoryId  string            `json:"category_id" gorm:"column:CategoryId"`
	Category    category.Category `gorm:"foreignkey:CategoryId"`
}

func (Item) TableName() string {
	return "Item"
}
