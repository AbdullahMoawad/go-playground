package controller

import (
	"github.com/google/uuid"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
)

func CreateSession(userId string) uuid.UUID {
	session := models.Session{}
	session.UserId = userId
	session.SessionId = uuid.New()
	serv.Conn().Create(&session)
	return session.SessionId
}

// @todo handle error
func CloseSession(SessionId string) {
	sessions := models.Session{}
	serv.Conn().Model(&sessions).Where("session_id = ?", SessionId).Unscoped().Delete(&sessions)
}
