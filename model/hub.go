package model

// Hub :
type Hub struct {
	ID          string     `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	GeoLocation [2]float64 `gorm:"column:geo_location"`
}

// TableName :
func (h Hub) TableName() string {
	return "hub"
}
