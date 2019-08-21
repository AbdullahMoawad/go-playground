package controller

import (
	"errors"
	"fmt"
	"github.com/sql-queries/common"
	"github.com/sql-queries/models"
	"github.com/sql-queries/server"
	"strings"
)

type User models.User

type UserLogin struct {
	Email 		string 	`gorm:"type:varchar(100);unique_index"`
	Password 	string
	IsActive 	bool
}

func (self *UserLogin) Format() *UserLogin {

	self.Email = strings.ToLower(self.Email)

	return self
}

func (self *User) FindByLogin(mail string) (error ,*User) {
	newUser := &User{}
	queryResult := server.Conn().Where(&User{Email: mail}).First(newUser)
	if queryResult.Error != nil {
		fmt.Println()
		return errors.New("Error while connecting to database "),nil
	}
	return nil ,newUser
}

func (self *UserLogin) ValidateLogin() (error, *User) {

	user := &User{}

	if self.Email != "" && self.Password != "" {
		_, user = user.FindByLogin(self.Email)
		if user == nil || user.Email == "" {
			return errors.New("Error login, user doesn't exist "), nil
		}else if self.Email != user.Email{
			return errors.New("Error login, Wrong email or password "), nil
		}
		password := common.CheckPasswordHash(self.Password,user.Password)
		if password == false {
			return errors.New("error login, Wrong email or password"), nil
		}
		if user.IsActive == false {
			return errors.New("please reactivate your account or call customer support"), nil
		}
	}
	user.Password = ""

	return nil, user
}