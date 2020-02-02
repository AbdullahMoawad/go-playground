package requests

import (
	"github.com/google/uuid"
	"real-estate/models"
)

type UserRequest struct {
	models.User
}

func NewUserRequest() *UserRequest {
	var user UserRequest
	user.Id = uuid.New().String()
	user.IsActive = true
	user.IsAdmin = false
	return &user
}
