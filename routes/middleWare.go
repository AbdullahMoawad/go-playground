package routes

import (
	"net/http"
	"real-estate/App"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/services"
)

func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("Sessionid")
		//@todo trim space
		//@check postgresql join
		if sessionId == "" {
			App.JsonLogger(w, 403, common.Login, nil)
			App.Logger(common.Login, sessionId)
			return
		}
		session := services.IsSessionExist(sessionId)
		if !session {
			App.JsonLogger(w, 440, common.SessionExpired, nil)
			App.Logger(common.SessionExpired, "error")
			return
		}
		f(w, r)
	}
}
