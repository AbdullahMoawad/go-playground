package routes

import (
	"encoding/json"
	"net/http"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/models"
	"real-estate/services"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("Sessionid")
		if sessionId == "" {
			 json.NewEncoder(w).Encode(models.Logger(401, common.Login, nil))
			return
		}
		session := services.IsSessionExist(sessionId)
		if !session {
			json.NewEncoder(w).Encode(models.Logger(401, common.SessionExpired, nil))
			return
		}
		f(w, r)
	}
}
