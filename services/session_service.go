package services

import (
	"real-estate/models"
	"real-estate/server"
)

func IsSessionExist(id string) bool {
	session := models.Session{}
	queryResult := server.Conn().Where(&models.Session{SessionId: id}).Find(&session)
	if queryResult.Error != nil {
		return false
	}
	return true
}
