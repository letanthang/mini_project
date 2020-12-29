package model

// Team :
type Team struct {
	ID       string `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	TeamType string `gorm:"column:team_type"`
	HubID    string `gorm:"column:hub_id"`
}

// TableName :
func (h Team) TableName() string {
	return "team"
}
