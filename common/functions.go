package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetSessionId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["Id"]
	return id
}
