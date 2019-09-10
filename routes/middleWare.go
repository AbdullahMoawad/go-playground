package routes

import (
	"encoding/json"
	"fmt"
	"github.com/sql-queries/models"
	"github.com/sql-queries/server"
	"net/http"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		sessionId := r.Header.Get("sessionId")
		_ = json.NewDecoder(r.Body).Decode(&user)
		db := server.Conn().Model(&user).Where("session_id = ?", sessionId).First(&user).Value
		dbUser ,err := json.Marshal(db)
		if err != nil{
			fmt.Println(err)
		}
		err = json.Unmarshal(dbUser, user)
		if (user.SessionId).String() == "00000000-0000-0000-0000-000000000000" {
			msg := "Please login"
			json.NewEncoder(w).Encode(msg)
			return
		} else {
			msg := "Already logged in"
			json.NewEncoder(w).Encode(msg)
			return
		}
		f(w, r)
	}
}
