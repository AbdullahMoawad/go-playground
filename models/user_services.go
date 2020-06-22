package models

import (
	"github.com/jinzhu/gorm"
	"property/common"
	"property/server"

	"strings"
)

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

func (self *User) UpdateFailedTries(userId string) interface{} {
	if queryResult := server.CreatePostgresDbConnection().Model(self).Where("id = ?", userId).Updates(map[string]interface{}{
		"FailedTriesCount":  self.FailedTriesCount,
		"LastFailedLoginAt": self.LastFailedLoginAt,
		"IsActive":          self.IsActive,
	}); queryResult.Error != nil {
		return queryResult.Error.Error()
	}
	return nil
}

func (self *User) FindByEmail(email string) *User {
	if queryResult := server.CreatePostgresDbConnection().Model(self).Where("email = ?", email).First(self); queryResult.Error != nil {
		return nil
	}
	return self
}

func (self *User) UpdateUser(email string) interface{} {
	if queryResult := server.CreatePostgresDbConnection().Model(&self).Where("email = ?", email).Updates(map[string]interface{}{
		"sessionId":        &self.SessionId,
		"FailedTriesCount": 0,
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

func GetPassword(id string) (error, string) {
	newUser := &User{}
	queryResult := server.CreatePostgresDbConnection().Model(&newUser).Where(&User{Id: id}).First(&newUser)
	if queryResult.Error != nil {
		return queryResult.Error, "error while getting user by id"
	} else {
		return nil, newUser.Password
	}
}

func (self *UserLogin) Format() *UserLogin {
	self.Email = strings.ToLower(self.Email)
	return self
}

func (self *User) Save() error {
	err := self.GetConnection().Save(self)
	return err.Error
}

func (self *UserLogin) ValidateLogin() (interface{}, *User) {
	var user User
	if self.Email != "" && self.Password != "" {

		user := user.FindByEmail(self.Email)

		if user.Email == "" {
			return common.UserNotFound, user
		}

		if !user.IsActive {
			return common.NotActiveUser, user
		}
		err, IsMatched := common.CheckPasswordHash(self.Password, user.Password)
		if !IsMatched {
			return err.Error(), user
		}
	}
	return nil, &user
}
