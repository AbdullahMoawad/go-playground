package services

import (
	"errors"
	"real-estate/models"
	"real-estate/server"
)

//func GetSessionId(r *http.Request) id string  {
//
//}
func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	user := &models.Session{}
	queryResult := server.Conn().Where(&User{SessionId: SessionID}).First(user)
	if queryResult.Error != nil {
		return errors.New("Error while connecting to database "), ""
	}

	return nil, user.UserId

}
