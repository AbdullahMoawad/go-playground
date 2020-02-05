package routes

import (
	"net/http"
	"real-estate/App"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/models"
)

func IsLogged(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controller := App.Controller{}
		var user models.User

		sessionId := r.Header.Get("Sessionid")
		//@todo trim space
		//@check postgresql join
		if sessionId == "" {
			App.Logger("error", common.Login, sessionId)
			controller.JsonLogger(w, 500, common.Login)
			return
		}
		session := user.IsSessionExist(sessionId)
		if !session {
			App.Logger("error", common.SessionExpired, "")
			controller.JsonLogger(w, 500, common.SessionExpired)
			return
		}
		f(w, r)
	}
}
