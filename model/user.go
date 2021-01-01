package model

// User :
type User struct {
	ID     string `json:"id" gorm:"column:id"`
	Role   string `json:"role" gorm:"column:role"`
	Email  string `json:"email" gorm:"column:email"`
	TeamID string `json:"team_id" gorm:"column:team_id"`
}

// TableName :
func (h User) TableName() string {
	return "user"
}
