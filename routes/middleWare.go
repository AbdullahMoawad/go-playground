package routes

import (
	"encoding/json"
	"log"
	"net/http"
	_ "real-estate/controller"
	"real-estate/models"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session := &models.Session{}
		session.SessionId = r.Header.Get("sessionId")
		if session.SessionId == "00000000-0000-0000-0000-000000000000" {
			json.NewEncoder(w).Encode(models.Logger(401,"Please login"))
			return
		}

		err, _ := findSession(session.SessionId)
		if err != nil{
			log.Println(err.Error())
			json.NewEncoder(w).Encode(models.Logger(401,"Please login,session expired"))
			return
		}
		f(w, r)
	}
}
