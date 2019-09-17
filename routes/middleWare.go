package routes

import (
	"encoding/json"
	"net"
	"net/http"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/models"
	"real-estate/services"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := &models.Session{}
		ip,_,_ := net.SplitHostPort(r.RemoteAddr)
		session.Ip = ip
		session.SessionId = r.Header.Get("sessionId")
		if session.SessionId == common.EmptySessionId {
			json.NewEncoder(w).Encode(models.Logger(401, common.Login, nil))
			return
		}
		err, _ :=  services.FindSession(session.SessionId)
		if err != nil {
			json.NewEncoder(w).Encode(models.Logger(401, common.SessionExpired, err))
			return
		}
		f(w, r)
	}
}
