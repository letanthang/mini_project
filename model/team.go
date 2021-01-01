package model

// Team :
type Team struct {
	ID       string `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	TeamType string `json:"team_type" gorm:"column:team_type"`
	HubID    string `json:"hub_id" gorm:"column:hub_id"`
}

// TableName :
func (h Team) TableName() string {
	return "team"
}
