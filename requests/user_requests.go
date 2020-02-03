package requests

import (
	"github.com/google/uuid"
	"real-estate/models"
)


func NewUserRequest() *models.User {
	var user models.User
	user.Id = uuid.New().String()
	user.IsActive = true
	user.IsAdmin = false
	return &user
}
