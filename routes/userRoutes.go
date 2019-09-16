package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/controller"
	"real-estate/uploader"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", IsLoggedin(controller.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/profile/{id}", IsLoggedin(controller.Profile)).Methods("GET")
	r.HandleFunc("/user/deactivate", IsLoggedin(controller.DeactivateUser)).Methods("POST")
	r.HandleFunc("/user/login", controller.Login).Methods("POST")
	r.HandleFunc("/user/logout", controller.Logout).Methods("DELETE")

	//Real Estate Routes
	r.HandleFunc("/estate", IsLoggedin(controller.CreateEstate)).Methods("POST")
	r.HandleFunc("/estate/{id}", IsLoggedin(controller.UpdateEstate)).Methods("PUT")
	r.HandleFunc("/estate/all", IsLoggedin(controller.ListEstates)).Methods("GET")
	r.HandleFunc("/estate/one/{id}", IsLoggedin(controller.OneEstate)).Methods("GET")
	r.HandleFunc("/estate/{id}", IsLoggedin(controller.DeleteEstate)).Methods("DELETE")

	//Category Routes
	r.HandleFunc("/category", IsLoggedin(controller.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/all", IsLoggedin(controller.ListCategories)).Methods("GET")
	r.HandleFunc("/category/one/{id}", IsLoggedin(controller.OneCategory)).Methods("GET")
	r.HandleFunc("/category/{id}", IsLoggedin(controller.DeleteCategory)).Methods("DELETE")

	// Upload file
	r.HandleFunc("/upload", IsLoggedin(uploader.UploadFile)).Methods("POST")

	http.ListenAndServe(":8000", r)
}
