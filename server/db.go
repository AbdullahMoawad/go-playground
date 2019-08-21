package server

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/sql-queries/models"
)
var db *gorm.DB

func Conn () *gorm.DB {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=abdullah dbname=queries password=root sslmode=disable" )
    if err != nil{
        panic(err)
    }
    db.AutoMigrate(models.User{})
    return db
}


