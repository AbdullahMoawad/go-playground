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
		"nickName":    user.NickName,
		"firsName":    user.FirstName,
		"lastName":    user.LastName,
		"password":    user.Password,
		"email":       user.Email,
		"address":     user.Address,
		"phoneNumber": user.PhoneNumber,
		"isActive":    user.IsActive	,});
	err != nil {
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
	//var user *models.User
	session, err := serv.Store.Get(r, "login")
	if err != nil{
		panic(err)
	}
	var userLogin *UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		panic(err)
	}

	err, user := userLogin.Format().ValidateLogin()
	if err != nil{
		_ = json.NewEncoder(w).Encode(err)

		return
	}


	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
	session.Values["authenticated"] = true
	if  err = session.Save(r, w); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

}

func Deactivate(w http.ResponseWriter, r *http.Request) {
	//var user *models.User
	session, err := serv.Store.Get(r, "login")
	if err != nil{
		panic(err)
	}
	var userLogin *UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		panic(err)
	}

	err, user := userLogin.Format().ValidateLogin()
	if err != nil{
		_ = json.NewEncoder(w).Encode(err)

		return
	}


	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
	session.Values["authenticated"] = true
	if  err = session.Save(r, w); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}

	if err := serv.Conn().Create(&user); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}