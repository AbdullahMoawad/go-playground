package controller

import (
	"property/App"
	"property/models"
	"property/server"
)

type SessionController struct {
	App.Controller
}

func CloseSession(SessionId string) {
	sessions := models.Session{}
	if queryResult := server.CreatePostgresDbConnection().Model(&sessions).Where("session_id = ?", SessionId).Unscoped().Delete(&sessions); queryResult.Error != nil {
		App.Logger("error", "Error create category", queryResult.Error.Error())
		return
	}
}
