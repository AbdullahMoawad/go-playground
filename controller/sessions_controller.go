package controller

import (
	"github.com/google/uuid"
	"real-estate/models"
	serv "real-estate/server"
)

func CreateSession(userId string) string {
	session := models.Session{}
	session.UserId = userId
	session.SessionId = uuid.New().String()
	serv.Conn().Create(&session)
	return session.SessionId
}

// @todo handle error
func CloseSession(SessionId string) {
	sessions := models.Session{}
	serv.Conn().Model(&sessions).Where("session_id = ?", SessionId).Unscoped().Delete(&sessions)
}
