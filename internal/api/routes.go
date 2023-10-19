package api

import (
	"app/internal/controller"
	"app/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func InitRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	r.HandleFunc("/users", userController.Create).Methods("POST")

	return r
}

func StartServer(db *gorm.DB) {
	r := InitRouter(db)
	port := ":8080"
	log.Println(fmt.Sprintf("App running successfully on port %s ...", port))
	http.ListenAndServe(port, r)
}
