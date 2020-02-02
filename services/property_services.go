package services

import (
	"github.com/jinzhu/gorm"
	"real-estate/models"
	"real-estate/server"
)

func FindAllProperties(userId string) *gorm.DB {
	var property []models.Property
	queryResult := server.Conn().Where("user_id = ?", userId).Find(&property)
	return queryResult
}
