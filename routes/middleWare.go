package routes

import (
	"net/http"
	"real-estate/App"
	"real-estate/common"
	_ "real-estate/controller"
	"real-estate/services"
)

type MiddlewareController struct {
	App.Controller
}

func IsLogged(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("Sessionid")
		//@todo trim space
		//@check postgresql join
		if sessionId == "" {
			App.Logger(common.Login, sessionId)
			return
		}
		session := services.IsSessionExist(sessionId)
		if !session {
			App.Logger(common.SessionExpired, "error")

			return
		}
		f(w, r)
	}
}
