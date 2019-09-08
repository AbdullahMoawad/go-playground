package controller

import (
	"encoding/json"
	"fmt"
	"github.com/sql-queries/common"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}
	user.Password = common.HashPassword(user.Password)

	// validateduser := user.Validate()
	if err := serv.Conn().Model(&user).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"nickName":     user.NickName,
		"firsName":     user.FirstName,
		"lastName":     user.LastName,
		"password":     user.Password,
		"email":        user.Email,
		"address":      user.Address,
		"phoneNumber":  user.PhoneNumber,
		"isAdmin":      user.IsAdmin,
		"isSuperAdmin": user.IsSuperAdmin,
		"isActive":     user.IsActive}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		fmt.Println(err)
	}
	newUser.Password = common.HashPassword(newUser.Password)
	if err := serv.Conn().Create(&newUser); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		panic(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userLogin *UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		panic(err)
	}
	err, user := userLogin.Format().ValidateLogin()
	if err != nil {
		panic(err)
	}
	user.SessionId = CreateSession(user.ID)
	if err := serv.Conn().Model(&user).Where("email = ?", userLogin.Email).Updates(map[string]interface{}{
		"session_id" :  user.SessionId,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessId := r.Header.Get("sessionId")
	CloseSession(sessId)
	_ = json.NewEncoder(w).Encode("logged out successfully ")
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
// @todo when check if logged in i will send session id in header if existe = detele , if not existe return error   not logged in
// @todo create proprties