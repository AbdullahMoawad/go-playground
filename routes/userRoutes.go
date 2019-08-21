package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sql-queries/controller"
	"net/http"
)

func Routes()  {
	r := mux.NewRouter()
	r.HandleFunc("/user", Logging(controller.CreateUser)).Methods("POST")
	r.HandleFunc("/update", Logging(controller.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/login", Logging(controller.Login)).Methods("POST")
	r.HandleFunc("/user/logout", Logging(controller.Logout)).Methods("PUT")
	fmt.Println(http.ListenAndServe(":8000", r))
}

