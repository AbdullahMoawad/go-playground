package models

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"real-estate/common"
	"real-estate/server"
	"strings"
)

type UserLogin struct {
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
	SessionId string `json:"sessionId"`
	IsActive  bool   `json:"isActive"`
}

func (self *User) Create() interface{} {
	if queryResult := server.CreatePostgresDbConnection().Create(&self); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return nil
}

func (self *User) Update(id string) interface{} {
	if queryResult := server.CreatePostgresDbConnection().Model(self).Where("id = ?", id).Updates(self); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return nil
}

func (self *User) UpdateUserSessionId(email string) interface{} {
	if queryResult := server.CreatePostgresDbConnection().Model(&self).Where("email = ?", email).Updates(map[string]interface{}{
		"sessionId": &self.SessionId,
	}); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return nil
}

func (self *User) Me() interface{} {
	queryResult := server.CreatePostgresDbConnection().Where("id = ?", self.Id).Find(&self)
	if queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	self.Password = ""
	self.SessionId = ""
	return self
}

func (self *User) GetConnection() *gorm.DB {
	return server.CreatePostgresDbConnection()
}

func (self *User) FindByEmail(mail string) (error, *User) {
	newUser := &User{}
	queryResult := server.CreatePostgresDbConnection().Where(&User{Email: mail}).First(newUser)
	if queryResult.Error != nil {
		return queryResult.Error, nil
	} else {
		return nil, newUser
	}
}

func GetPassword(id string) (error, string) {
	newUser := &User{}
	queryResult := server.CreatePostgresDbConnection().Model(&newUser).Where(&User{Id: id}).First(&newUser)
	if queryResult.Error != nil {
		return queryResult.Error, "error while getting user by id"
	} else {
		return nil, newUser.Password
	}
}

//func (self *User) Create() error {
//	err := self.GetConnection().Create(self)
//	return err.Error
//}

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
		_, IsMatched := common.CheckPasswordHash(self.Password, user.Password)
		if IsMatched == false {
			return common.LoginFailed, nil
		}
		if user.IsActive == false {
			return common.NotActiveUser, nil
		}
	} else {
		return common.UserNotFound, nil
	}
	user.Password = ""
	return "", user
}

func GetCurrentUserIdFromHeaders(SessionID string) (error, string) {
	session := Session{}
	queryResult := server.CreatePostgresDbConnection().Where(&Session{SessionId: SessionID}).First(&session)
	return queryResult.Error, session.UserId
}

func GetCurrentUserIdByEmail(Email string) string {
	user := User{}
	if queryResult := server.CreatePostgresDbConnection().Where(&User{Email: Email}).First(&user); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return user.Id
}

func GetCurrentUserId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}
