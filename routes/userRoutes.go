package routes

import (
	"github.com/gorilla/mux"
	"github.com/sql-queries/controller"
	"net/http"
)

func Routes()  {
	r := mux.NewRouter()
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/me", controller.ProfielOfUser).Methods("PUT")
	r.HandleFunc("/user/deactivate", controller.DeactivateUser).Methods("POST")
	r.HandleFunc("/user/login", controller.Login).Methods("POST")
	r.HandleFunc("/user/logout", controller.Logout).Methods("DELETE")

	//Real Estate Routes
	r.HandleFunc("/estate", controller.CreateRealEstate).Methods("POST")
	r.HandleFunc("/estate", controller.UpdateRealEstate).Methods("PUT")
	r.HandleFunc("/estate/my", controller.MyRealEstates).Methods("GET")
	r.HandleFunc("/estate/list", controller.ListRealEstates).Methods("GET")
	r.HandleFunc("/estate/all", controller.ListAllRealEstates).Methods("GET")
	r.HandleFunc("/estate/{id}", controller.DeleteRealEstates).Methods("DELETE")

	//Category Routes
	r.HandleFunc("/category", controller.CreateCategory).Methods("POST")
	r.HandleFunc("/category", controller.UpdateCategory).Methods("PUT")
	r.HandleFunc("/category/list", controller.ListCategory).Methods("GET")
	r.HandleFunc("/category/{id}", controller.DeleteCategory).Methods("DELETE")

	_ = http.ListenAndServe(":8000", r)
}