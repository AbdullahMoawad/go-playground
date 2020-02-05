package models

import (
	"github.com/google/uuid"
	"net/http"
	"real-estate/server"
)

func CreateSession(userId string) (error, string) {
	var session Session
	session.UserId = userId
	session.SessionId = uuid.New().String()
	if queryResult := server.CreatePostgresDbConnection().Create(&session); queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, session.SessionId
}

func IsSessionExist(id string) bool {
	session := Session{}
	queryResult := server.CreatePostgresDbConnection().Where(&Session{SessionId: id}).Find(&session)
	if queryResult.Error != nil {
		return false
	}
	return true
}
func GetCurrentSessionId(r *http.Request) string {
	sessionId := GetCurrentSessionId(r)
	return sessionId
}
