package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	"real-estate/requests"
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
	sms.SendSms("welcome to our web site welcome message test :)","13393371991",newUser.PhoneNumber)
	self.Json(w, newUser, common.StatusOK)
}

func (self UserController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userRequest := requests.NewUserRequest()

	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	if err := serv.Conn().Model(&userRequest).Where("id = ?", id).Updates(map[string]interface{}{
		"nickName": userRequest.NickName,
		"firsName": userRequest.FirstName,
		"lastName": userRequest.LastName,
		"address":  userRequest.Address,
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
	id := common.GetId(r)

	queryResult := serv.Conn().Model(&user).Where("id = ?", id).First(&user)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, common.ProfileError, queryResult)
		self.Logger(common.ProfileError, "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self UserController) Deactivate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		self.JsonLogger(w, 500, common.DecodingError, err)
		self.Logger(common.DecodingError, "error")
		return
	}

	if err := serv.Conn().Model(&user).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
		"isActive": false}); err != nil {
		self.JsonLogger(w, 500, "error while deactivating user", err)
		self.Logger("error while deactivating user", "error")
		return
	}

	self.Json(w, user, common.StatusOK)
}
