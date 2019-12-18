package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newUser := models.NewUser()

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, common.DecodingError, err))
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
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, common.DecodingError, err))
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
	defer r.Body.Close()

	var userLogin *services.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	err, user := userLogin.Format().ValidateLogin()
	if err != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	// set session id
	user.SessionId = CreateSession(user.Id)

	if err := serv.Conn().Model(&user).Where("email = ?", userLogin.Email).Updates(map[string]interface{}{"session_id": user.SessionId,}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessId := r.Header.Get("sessionId")
	CloseSession(sessId)
	json.NewEncoder(w).Encode("Logged out successfully ")
}

func Profile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	id := common.GetId(r)

	queryResult := serv.Conn().Model(&user).Where("id = ?", id).First(&user)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, common.ProfileError, queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func DeactivateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(models.Logger(500, common.DecodingError, err))
		return
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
