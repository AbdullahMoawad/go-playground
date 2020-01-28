package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/controller"
	"real-estate/uploader"
)

func Routes() {
	r := mux.NewRouter()
	userController := controller.UserController{}

	r.HandleFunc("/user", userController.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", IsLoggedin(userController.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/me/{id}", IsLoggedin(userController.Profile)).Methods("GET")
	r.HandleFunc("/user/deactivate", IsLoggedin(userController.DeactivateUser)).Methods("POST")
	r.HandleFunc("/user/login", userController.Login).Methods("POST")
	//@todo use http.MethodX
	r.HandleFunc("/user/logout", userController.Logout).Methods(http.MethodDelete)

	//Real Estate Routes

	RealEstateController := controller.RealEstateController{}
	r.HandleFunc("/estate", IsLoggedin(RealEstateController.CreateRealEstate)).Methods("POST")
	r.HandleFunc("/estate/{id}", IsLoggedin(RealEstateController.UpdateRealEstate)).Methods("PUT")
	r.HandleFunc("/estate/all", IsLoggedin(RealEstateController.ListRealEstate)).Methods("GET")
	r.HandleFunc("/estate/one/{id}", IsLoggedin(RealEstateController.OneRealEstate)).Methods("GET")
	r.HandleFunc("/estate/{id}", IsLoggedin(RealEstateController.DeleteRealEstate)).Methods("DELETE")

	//Category Routes

	CategoryController := controller.CategoryController{}
	r.HandleFunc("/category", IsLoggedin(CategoryController.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/all", IsLoggedin(CategoryController.ListCategories)).Methods("GET")
	r.HandleFunc("/category/{id}", IsLoggedin(CategoryController.OneCategory)).Methods("GET")
	r.HandleFunc("/category/{id}", IsLoggedin(CategoryController.DeleteCategory)).Methods("DELETE")

	// Upload file
	r.HandleFunc("/upload", IsLoggedin(uploader.UploadFile)).Methods("POST")

	http.ListenAndServe(":8000", r)
}
