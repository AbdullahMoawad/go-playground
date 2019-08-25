package controller

import (
	"encoding/json"
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
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}

	user.Password = common.HashPassword(user.Password)
	if err := serv.Conn().Create(&user); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	_, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var userLogin *UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		panic(err)
	}

	err, user := userLogin.Format().ValidateLogin()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
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
