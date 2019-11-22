package services

import (
	"real-estate/models"
	"real-estate/server"
)

func IsExist(name string) bool {
	category := models.Category{}
	queryResult := server.Conn().Where(&models.Category{Name: name}).Find(&category)
	if queryResult.Error != nil {
		return false
	}
	return true
}
