package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
	"real-estate/sms"
)

type UserController struct {
	App.Controller
}

func (self UserController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newUser := models.NewUser()

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	newUser.Password = common.HashPassword(newUser.Password)
	if err := serv.Conn().Create(newUser); err.Error != nil {
		self.JsonLogger(w, 500, "Error creating user", err)
		self.Logger("Error creating user", "error")
		return
	}

	sms.SendSms("welcome to our web site welcome message test :)", "13393371991", newUser.PhoneNumber)
	self.Json(w, newUser, common.StatusOK)
}

func (self UserController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var User *models.User

	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	if err := serv.Conn().Model(&User).Where("id = ?", id).Updates(map[string]interface{}{
		"nickName": User.NickName,
		"firsName": User.FirstName,
		"lastName": User.LastName,
		"address":  User.Address,
	}); err != nil {
		self.Json(w, err, common.StatusOK)
		return
	}
}

// @todo review the code
func (self UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userLogin *services.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	err, user := userLogin.Format().ValidateLogin()
	if err != "" {
		self.Json(w, err, common.StatusOK)
		return
	}
	// set session id
	user.SessionId = CreateSession(user.Id)

	if err := serv.Conn().Model(&user).Where("email = ?", userLogin.Email).Updates(map[string]interface{}{"session_id": user.SessionId}); err != nil {
		self.Json(w, err, common.StatusOK)
		return
	}
	self.Json(w, user, common.StatusOK)
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
	userId := common.GetId(r)

	queryResult := serv.Conn().Model(&user).Where("id = ?", userId).First(&user)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, common.ProfileError, queryResult)
		self.Logger(common.ProfileError, "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self UserController) Deactivate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user *models.User
	var session models.Session
	tx := serv.Conn().Begin()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	userId := common.GetId(r)
	if userId == "" {
		self.JsonLogger(w, 400, common.EmptyUserId, nil)
		self.Logger(common.EmptyUserId, "error")
		return
	}

	err, userPassword := services.FindById(userId)
	if err != nil {
		self.JsonLogger(w, 400, common.UserNotFound, err)
		self.Logger(common.UserNotFound, "error")
		return
	}

	if !common.CheckPasswordHash(user.Password, userPassword) {
		self.JsonLogger(w, 400, common.WorngPassword, err)
		self.Logger(common.WorngPassword, "error")
		return
	}

	if err := tx.Model(&user).Where("id = ?", userId).Updates(map[string]interface{}{
		"isActive": false}); err.Error != nil {
		tx.Rollback()
		self.JsonLogger(w, 400, "error while deactivating user", err.Error)
		self.Logger("error while deactivating user", "error")
		return
	}

	if err := tx.Where("user_idd = ?",userId ).Unscoped().Delete(&session); err.Error != nil {
		tx.Rollback()
		self.JsonLogger(w, 400, "error while deleting session", err.Error)
		self.Logger(" error while deleting session", "error")
	}
	tx.Commit()

	self.Json(w, "user deactivated successfully", common.StatusOK)
}
