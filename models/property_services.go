package models

import (
	"github.com/jinzhu/gorm"
	"property/server"
)

func FindAllProperties(userId string) *gorm.DB {
	var property []Property
	queryResult := server.CreatePostgresDbConnection().Where("user_id = ?", userId).Find(&property)
	return queryResult
}
