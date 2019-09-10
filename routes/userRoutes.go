package routes

import (
	"github.com/gorilla/mux"
	"github.com/sql-queries/controller"
	"net/http"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/user", IsLoggedIn(controller.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}", IsLoggedIn(controller.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/profile/{id}", IsLoggedIn(controller.Profile)).Methods("GET")
	r.HandleFunc("/user/deactivate", IsLoggedIn(controller.DeactivateUser)).Methods("POST")
	r.HandleFunc("/user/login", controller.Login).Methods("POST")
	r.HandleFunc("/user/logout", controller.Logout).Methods("DELETE")

	//Real Estate Routes
	r.HandleFunc("/estate", IsLoggedIn(controller.CreateEstate)).Methods("POST")
	r.HandleFunc("/estate/{estateId}", IsLoggedIn(controller.UpdateEstate)).Methods("PUT")
	//r.HandleFunc("/estate/my", controller.MyRealEstates).Methods("GET")
	//r.HandleFunc("/estate/list", controller.ListRealEstates).Methods("GET")
	//r.HandleFunc("/estate/all", controller.ListAllRealEstates).Methods("GET")
	//r.HandleFunc("/estate/{id}", controller.DeleteRealEstates).Methods("DELETE")

	//Category Routes
	//r.HandleFunc("/category", controller.CreateCategory).Methods("POST")
	//r.HandleFunc("/category", controller.UpdateCategory).Methods("PUT")
	//r.HandleFunc("/category/list", controller.ListCategory).Methods("GET")
	//r.HandleFunc("/category/{id}", controller.DeleteCategory).Methods("DELETE")
	http.ListenAndServe(":8000", r)
}
