package routes

import (
	"encoding/json"
	"github.com/sql-queries/models"
	"github.com/sql-queries/server"
	"net/http"
)

func IsLoggedIn(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user := models.User{}
		sessionId := r.Header.Get("sessionId")
		if sessionId == "" {
			msg := "Session Error"
			json.NewEncoder(w).Encode(msg)
			return
		}
		userSession := server.Conn().Model(&user).Where("session_id = ?", sessionId).First(&user)
		if userSession.Error != nil {
			msg := "Please Login"
			json.NewEncoder(w).Encode(msg)
			return
		}

		if user.SessionId == "00000000-0000-0000-0000-000000000000" {
			msg := "Please login"
			json.NewEncoder(w).Encode(msg)
			return
		}
		f(w, r)
	}
}
