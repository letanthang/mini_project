package model

// Hub :
type Hub struct {
	ID           string  `json:"id" gorm:"column:id"`
	Name         string  `json:"name" gorm:"column:name"`
	LocationLat  float64 `json:"location_lat" gorm:"column:location_lat"`
	LocationLong float64 `json:"location_long" gorm:"column:location_long"`
}

// TableName :
func (h Hub) TableName() string {
	return "hub"
}

// IsExisted :
func (h Hub) IsExisted() bool {
	if h.ID == "" {
		return false
	}
	return true
}
