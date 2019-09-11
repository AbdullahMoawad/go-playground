package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-estate/models"
	"real-estate/server"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		sessionId := r.Header.Get("sessionId")
		db := server.Conn().Model(&user).Where("session_id = ?", sessionId).First(&user).Value
		dbUser, err := json.Marshal(db)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(dbUser, user)
		if (user.SessionId).String() == "00000000-0000-0000-0000-000000000000" {
			msg := "Please login"
			json.NewEncoder(w).Encode(msg)
			return
		}
		f(w, r)
	}
}
