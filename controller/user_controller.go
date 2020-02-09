package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/common/helpers"
	"real-estate/models"
	"real-estate/server"
	"real-estate/sms"
	"time"
)

type UserController struct {
	App.Controller
}

func (self UserController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newUser := models.NewUser()

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		self.JsonLogger(w, 500, common.DecodingError)
		self.Logger("error", common.DecodingError, err)
		return
	}

	newUser.Password = common.HashPassword(newUser.Password)

	if queryResult := newUser.Create(); queryResult != nil {
		self.JsonLogger(w, 500, "Error creating user")
		self.Logger("error", common.DatabaseOperationFailed, queryResult)
		return
	}

	sms.SendSms("welcome to our web site welcome message test :)", newUser.PhoneNumber)
	self.Json(w, newUser, common.StatusOK)
}

func (self UserController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var User models.User
	userId := helpers.GetCurrentUserId(r)

	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		self.JsonLogger(w, 500, common.DecodingError)
		self.Logger("error", common.DecodingError, err)
		return
	}

	err := User.Update(userId)
	if err != nil {
		self.JsonLogger(w, 404, "No property found ..")
		self.Logger("error", common.DatabaseOperationFailed, err)
	}
	self.Json(w, User, common.StatusOK)
}

// @todo revamp the code
func (self UserController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var LoginRequest *models.UserLogin
	var User models.User

	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		self.JsonLogger(w, 500, common.DecodingError)
		self.Logger("error", common.DecodingError, err)
		return
	}

	err, user := LoginRequest.Format().ValidateLogin()
	if err != nil {
		if user != nil && user.IsActive {

			user.FailedTriesCount = user.FailedTriesCount + 1
			user.LastFailedLoginAt = time.Now()

			if user.FailedTriesCount >= 5 {
				user.IsActive = false
			}

			if err := user.UpdateFailedTries(user.Id); err != nil {
				self.JsonLogger(w, 500, "saving lock tires error")
				self.Logger("error", "saving lock tires error", err)
				return
			}
		}
		self.JsonLogger(w, 500, (err).(string))
		self.Logger("error", common.UserFormatingAndValidatingError, err)
		return
	}

	userId := helpers.GetCurrentUserIdByEmail(LoginRequest.Email)

	errMsg, sessionId := models.CreateSession(userId)
	if errMsg != nil {
		self.JsonLogger(w, 500, "error while creating session"+errMsg.Error())
		self.Logger("error", common.DatabaseOperationFailed, errMsg)
		return
	}

	user.SessionId = sessionId
	User.SessionId = user.SessionId

	// reset failed tries and update sessionId
	updateSessionId := User.UpdateUser(LoginRequest.Email)
	if updateSessionId != nil {
		self.JsonLogger(w, 400, "error while updating user Session")
		self.Logger("error", common.DatabaseOperationFailed, updateSessionId)
		return
	}

	self.Json(w, user, common.StatusOK)
}

func (self UserController) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sessionId := models.GetCurrentSessionId(r)
	CloseSession(sessionId)
	json.NewEncoder(w).Encode("Logged out successfully ")
}

func (self UserController) Profile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	user.Id = helpers.GetCurrentUserId(r)

	query := user.Me()

	self.Json(w, query, common.StatusOK)
}

func (self UserController) Deactivate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		self.JsonLogger(w, 500, common.DecodingError)
		self.Logger("error", common.DecodingError, err)
		return
	}

	userId := helpers.GetCurrentUserId(r)
	if userId == "" {
		self.JsonLogger(w, 400, common.EmptyUserId)
		self.Logger("error", common.EmptyUserId, nil)
		return
	}

	err, userPassword := models.GetPassword(userId)
	if err != nil {
		self.JsonLogger(w, 400, common.UserNotFound)
		self.Logger("error", common.UserNotFound, err)
		return
	}

	errMsg, IsMatched := common.CheckPasswordHash(user.Password, userPassword)
	if !IsMatched {
		self.JsonLogger(w, 400, common.WorngPassword)
		self.Logger("error", common.WorngPassword, errMsg)
		return
	}

	err, _ = user.Deactivate(userId)
	if err != nil {
		self.JsonLogger(w, 404, "Error deactivate user ..")
		self.Logger("error", common.DatabaseOperationFailed, err)
	}

	if queryResult := server.CreatePostgresDbConnection().Model(&user).Where("id = ?", userId).Updates(map[string]interface{}{
		"isActive": false}); queryResult.Error != nil {
		self.JsonLogger(w, 400, "error while deactivating user")
		self.Logger("error", "error while deactivating user", queryResult.Error.Error())
		return
	}

	err, _ = user.DeleteSession(userId)
	if err != nil {
		self.JsonLogger(w, 404, "Error delete session ..")
		self.Logger("error", common.DatabaseOperationFailed, err)
	}

	self.Json(w, "user deactivated successfully", common.StatusOK)
}
