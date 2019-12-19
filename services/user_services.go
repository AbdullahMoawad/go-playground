package services

import (
	"real-estate/common"
	"real-estate/models"
	"real-estate/server"
	"strings"
)

type User models.User

type UserLogin struct {
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
	SessionId string `json:"sessionId"`
	IsActive  bool   `json:"isActive"`
}

type Deactivate struct {
	Email    string `gorm:"type:varchar(100);unique_index"`
	IsActive bool
}

func (self *UserLogin) Format() *UserLogin {
	self.Email = strings.ToLower(self.Email)
	return self
}

func (self *UserLogin) ValidateLogin() (string, *User) {
	user := &User{}

	if self.Email != "" && self.Password != "" {
		_, user = user.FindByEmail(self.Email)
		if user == nil || user.Email == "" {
			return common.UserNotFound, nil
		} else if self.Email != user.Email {
			return common.LoginFailed, nil
		}
		password := common.CheckPasswordHash(self.Password, user.Password)
		if password == false {
			return common.LoginFailed, nil
		}
		if user.IsActive == false {
			return common.EmptyLoginFields, nil
		}
	} else {
		return common.UserNotFound, nil
	}
	user.Password = ""
	return "", user
}

func (self *User) FindByEmail(mail string) (error, *User) {
	newUser := &User{}
	queryResult := server.Conn().Where(&User{Email: mail}).First(newUser)
	if queryResult.Error != nil {
		return queryResult.Error, nil
	} else {
		return nil, newUser
	}
}

//func (self *User) GetCurrentUserFromHeaders(SessionID string) (error, string) {
//	user := &User{}
//	queryResult := server.Conn().Where(&User{SessionId: SessionID}).First(user)
//	if queryResult.Error != nil {
//		fmt.Println()
//		return queryResult.Error, ""
//	} else {
//		return nil, user.Email
//	}
//}
