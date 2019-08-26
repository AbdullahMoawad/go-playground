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
		sessId := r.Header.Get("sessionId")
		json.NewDecoder(r.Body).Decode(&user)
		coo := server.Conn().Model(&user).Where("session_id = ?", sessId).First(&user).Value
		neww ,err := json.Marshal(coo)
		if err != nil{
			fmt.Println(err)
		}

		err = json.Unmarshal(neww, user)

		if (user.SessionId).String() == "00000000-0000-0000-0000-000000000000" {
			fmt.Println("please login frist")
		} else {
			fmt.Println("already logged in")
		}
		f(w, r)
	}
}

