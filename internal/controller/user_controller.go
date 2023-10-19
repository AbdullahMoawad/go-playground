package controller

import (
	"app/App"
	"app/common"
	"app/internal/models"
	"app/internal/service"
	"encoding/json"
	"net/http"
)

type UserController struct {
	App.Controller
	userService *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{userService: s}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		c.respondWithError(w, http.StatusBadRequest, common.DecodingError)
		return
	}

	newUser.Password = common.HashPassword(newUser.Password)

	if err := c.userService.CreateUser(&newUser); err != nil {
		c.respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	c.respondWithJSON(w, http.StatusCreated, newUser)
}

func (c *UserController) respondWithError(w http.ResponseWriter, code int, message string) {
	c.JsonLogger(w, code, message)
	c.Logger("error", message, nil)
}

func (c *UserController) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	c.Json(w, payload, common.StatusOK)
}
