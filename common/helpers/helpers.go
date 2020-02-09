package helpers

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/models"
	"real-estate/server"
)

func GetCurrentUserId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetCurrentPropertyId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	session := models.Session{}
	queryResult := server.CreatePostgresDbConnection().Where(&models.Session{SessionId: SessionID}).First(&session)
	return queryResult.Error, session.UserId
}

func GetCurrentUserIdByEmail(Email string) string {
	user := models.User{}
	if queryResult := server.CreatePostgresDbConnection().Where(&models.User{Email: Email}).First(&user); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return user.Id
}
