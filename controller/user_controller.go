package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sql-queries/common"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	newUser.Password = common.HashPassword(newUser.Password)
	if err := serv.Conn().Create(&newUser); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user *models.User
	params := mux.Vars(r)
	id := params["id"]
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := serv.Conn().Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"nickName": user.NickName,
		"firsName": user.FirstName,
		"lastName": user.LastName,
		"address":  user.Address,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin *UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	err, user := userLogin.Format().ValidateLogin()
	if err != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	user.SessionId = CreateSession(user.Id)
	if err := serv.Conn().Model(&user).Where("email = ?", userLogin.Email).Updates(map[string]interface{}{
		"session_id": user.SessionId,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessId := r.Header.Get("sessionId")
	CloseSession(sessId)
	json.NewEncoder(w).Encode("Logged out successfully ")
}

func Profile(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	var model *models.User
	user := model
	params := mux.Vars(r)
	id := params["id"]
 	queryResult := serv.Conn().Model(&user).Where("id = ?", id).First(user)
	if queryResult.Error != nil {
		fmt.Println()
		return
	}
	if err := json.NewEncoder(w).Encode(queryResult); err != nil {
		panic(err)
	}
}

func DeactivateUser(w http.ResponseWriter, r *http.Request) {

	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}

	if err := serv.Conn().Model(&user).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
		"isActive": false}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}