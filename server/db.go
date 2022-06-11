package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbConfig struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

func CreatePostgresDbConnection() *gorm.DB {
	dbConfig := DbConfig{Host: "127.0.0.1", Port: "5438", Username: "root", Password: "root", DbName: "app"}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.DbName, dbConfig.Password)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Can't connect to postgres dbs", err)
		return nil
	}
	return db
}
