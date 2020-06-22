package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"property/controller"
)

func Routes() {
	r := mux.NewRouter()
	user := controller.UserController{}
	category := controller.CategoryController{}
	property := controller.PropertyController{}

	r.HandleFunc("/user", user.Create).Methods("POST")
	r.HandleFunc("/user/{id}", IsLogged(user.Update)).Methods("PUT")
	r.HandleFunc("/user/me/{id}", IsLogged(user.Profile)).Methods("GET")
	r.HandleFunc("/user/deactivate/{id}", IsLogged(user.Deactivate)).Methods("POST")
	r.HandleFunc("/user/login", user.Login).Methods("POST")
	r.HandleFunc("/user/logout", user.Logout).Methods("DELETE")

	//Real Estate Routes
	r.HandleFunc("/estate", IsLogged(property.Create)).Methods("POST")
	r.HandleFunc("/estate/{id}", IsLogged(property.Update)).Methods("PUT")
	r.HandleFunc("/estate/all", IsLogged(property.List)).Methods("GET")
	r.HandleFunc("/estate/one/{id}", IsLogged(property.One)).Methods("GET")
	r.HandleFunc("/estate/{id}", IsLogged(property.Delete)).Methods("DELETE")

	//Category Routes
	r.HandleFunc("/category", IsLogged(category.Create)).Methods("POST")
	r.HandleFunc("/category/all", IsLogged(category.List)).Methods("GET")
	r.HandleFunc("/category/{id}", IsLogged(category.One)).Methods("GET")
	r.HandleFunc("/category/{id}", IsLogged(category.Delete)).Methods("DELETE")

	// Upload file
	//r.HandleFunc("/upload", IsLogged(uploader.UploadFile)).Methods("POST")
	log.Println("App running successfully ..:)")

	http.ListenAndServe(":8000", r)
}
