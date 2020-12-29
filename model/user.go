package model

// User :
type User struct {
	ID     string `gorm:"column:id"`
	Role   string `gorm:"column:name"`
	Email  string `gorm:"column:email"`
	TeamID string `gorm:"column:team_id"`
}

// TableName :
func (h User) TableName() string {
	return "user"
}
