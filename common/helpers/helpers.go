package helpers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetCurrentUserId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func GetCurrentPropertyId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}
