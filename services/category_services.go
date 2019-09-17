package services

import (
	"real-estate/models"
	"real-estate/server"
)

func IsExist(name string) bool {
	category := models.Category{}
	quereyResult := server.Conn().Where(&models.Category{Name: name}).Find(&category)
	if quereyResult.Error != nil {
		return false
	}
	return true
}
