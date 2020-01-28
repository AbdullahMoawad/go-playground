package common

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/models"
	"real-estate/server"
)



func GetId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetSessionId(r *http.Request) string {
	sessionId := r.Header.Get("Sessionid")
	return sessionId
}

func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	session := models.Session{}
	queryResult := server.Conn().Where(&models.Session{SessionId: SessionID}).First(&session)
	return queryResult.Error, session.UserId
}
