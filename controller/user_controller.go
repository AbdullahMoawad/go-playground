package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
	"reflect"
)

type UserController struct{}

func (self UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newUser := models.NewUser()

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {

		App.JsonLogger(w, 500, common.DecodingError, err)
		App.Logger(common.DecodingError, "error")
		return
	}

	newUser.Password = common.HashPassword(newUser.Password)

	if err := serv.Conn().Create(&newUser); err.Error != nil {

		zap := reflect.ValueOf(err.Error)
		fmt.Println(zap,"--=-=-=---")

		App.JsonLogger(w, 500, "Error creating user", zap)
		App.Logger("Error creating user", "error")
		return
	}

	App.Json(w, newUser, common.StatusOK)
}

func (self UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user *models.User
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		App.JsonLogger(w, 500, common.DecodingError, err)
		App.Logger(common.DecodingError, "error")
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

func (self UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userLogin *services.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		App.JsonLogger(w, 500, common.DecodingError, err)
		App.Logger(common.DecodingError, "error")
		return
	}
	err, user := userLogin.Format().ValidateLogin()
	if err != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	// set session id
	user.SessionId = CreateSession(user.Id)

	if err := serv.Conn().Model(&user).Where("email = ?", userLogin.Email).Updates(map[string]interface{}{"session_id": user.SessionId}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (self UserController) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessId := r.Header.Get("sessionId")
	CloseSession(sessId)
	json.NewEncoder(w).Encode("Logged out successfully ")
}

func (self UserController) Profile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	id := common.GetId(r)

	queryResult := serv.Conn().Model(&user).Where("id = ?", id).First(&user)
	if queryResult.Error != nil {
		App.JsonLogger(w, 500, common.ProfileError, queryResult)
		App.Logger(common.ProfileError, "error")
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func (self UserController) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		App.JsonLogger(w, 500, common.DecodingError, err)
		App.Logger(common.DecodingError, "error")
		return
	}

	if err := serv.Conn().Model(&user).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
		"isActive": false}); err != nil {
		App.JsonLogger(w, 500, "error while deactivating user", err)
		App.Logger("error while deactivating user", "error")
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}


