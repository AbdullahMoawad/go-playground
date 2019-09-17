package services

import (
	"real-estate/models"
	"real-estate/server"
)

func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	session := models.Session{}
	queryResult := server.Conn().Where(&models.Session{SessionId: SessionID}).First(&session)
	if queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, session.UserId
}
