package services

import (
	"real-estate/models"
	"real-estate/server"
)

type Session models.Session

func FindSession(sessionId string) (error, *Session) {
	session := &Session{}
	queryResult := server.Conn().Where(&models.Session{SessionId: sessionId}).First(&session)
	if queryResult.Error != nil {
		return queryResult.Error, nil
	}
	return nil, session
}
