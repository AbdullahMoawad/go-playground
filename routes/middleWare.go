package routes

import (
	"net/http"
	"real-estate/App"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/services"
)


func IsLogged(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controller := App.Controller{}
		sessionId := r.Header.Get("Sessionid")
		//@todo trim space
		//@check postgresql join
		if sessionId == "" {
			App.Logger(common.Login, sessionId)
			controller.JsonLogger(w, 500, common.Login, nil)
			return
		}
		session := services.IsSessionExist(sessionId)
		if !session {
			App.Logger(common.SessionExpired, "error")
			controller.JsonLogger(w, 500, common.SessionExpired, nil)
			return
		}
		f(w, r)
	}
}
