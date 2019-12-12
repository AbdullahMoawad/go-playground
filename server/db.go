package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Conn() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=macbookpro dbname=realestate password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
