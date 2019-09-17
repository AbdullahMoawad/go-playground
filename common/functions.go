package common

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/models"
	"real-estate/server"
)
type Session models.Session

func GetId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetSessionId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["Id"]
	return id
}

func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	session := models.Session{}
	queryResult := server.Conn().Where(&models.Session{SessionId: SessionID}).First(&session)
	if queryResult.Error != nil {
		return queryResult.Error, ""
	}
	return nil, session.UserId
}
