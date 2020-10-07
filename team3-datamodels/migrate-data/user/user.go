package user

type User struct {
	ID    string `json:"id" gorm:"column:UserId;PRIMARY KEY"`
	Name  string `json:"name" gorm:"column:"Name"`
	Email string `json:"email" gorm:"column:"Email"`
}

func (User) TableName() string {
	return "User"
}
