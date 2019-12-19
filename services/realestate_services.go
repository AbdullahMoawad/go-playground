package services

import (
	"github.com/jinzhu/gorm"
	"real-estate/models"
	"real-estate/server"
)

func FindAllEstates(userId string) *gorm.DB {
	var estates []models.RealEstate
	queryResult := server.Conn().Where("user_id = ?", userId).Find(&estates)
	return queryResult
}
