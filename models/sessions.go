package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Session struct {
	gorm.Model
	UserId    string
	SessionId string
	Ip        string
}
