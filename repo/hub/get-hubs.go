package hub

import (
	"app/model"
	"app/repo"
	"fmt"
)

func GetHubs() (*[]model.Hub, error) {
	var hubs []model.Hub
	if dbc := repo.GetDB().Find(&hubs); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &hubs, nil
}
