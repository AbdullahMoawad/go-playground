package routes

import (
	"net/http"
	"real-estate/App"
	_ "real-estate/controller"
	"real-estate/services"
)

type Controller struct {
	App.Controller
}

func IsLogged(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("Sessionid")
		//@todo trim space
		//@check postgresql join
		if sessionId == "" {
			//App.JsonLogger(w, 403, common.Login, nil)
			//self.Logger(common.Login, sessionId)
			return
		}
		session := services.IsSessionExist(sessionId)
		if !session {
			//self.JsonLogger(w, 440, common.SessionExpired, nil)
			//self.Logger(common.SessionExpired, "error")
			return
		}
		f(w, r)
	}
}
