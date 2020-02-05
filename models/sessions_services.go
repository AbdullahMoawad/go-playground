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

func (self *User) DeleteSession(userId string) (error, string) {
	var session Session
	if queryResult := server.CreatePostgresDbConnection().Where("user_id = ?", userId).Unscoped().Delete(&session); queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, "session deleted successfully"
}

func (self *User) Deactivate(userId string) (error, string) {
	var session Session
	if queryResult := server.CreatePostgresDbConnection().Model(session).Where("user_id = ?", userId).Update(map[string]interface{}{
		"isActive": false}); queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, "session deleted successfully"
}

func (self *User)IsSessionExist(id string) bool {
	session := Session{}
	queryResult := server.CreatePostgresDbConnection().Where(&Session{SessionId: id}).Find(&session)
	if queryResult.Error != nil {
		return false
	}
	return true
}

func GetCurrentSessionId(r *http.Request) string {
	sessionId := r.Header.Get("sessionId")
	return sessionId
}