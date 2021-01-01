package model

// Hub :
type Hub struct {
	ID          string     `json:"id" gorm:"column:id"`
	Name        string     `json:"name" gorm:"column:name"`
	GeoLocation [2]float64 `json:"geo_location" gorm:"column:geo_location"`
}

// TableName :
func (h Hub) TableName() string {
	return "hub"
}
