package main

import (
	"github.com/sql-queries/ops"
	setupRoutes "github.com/sql-queries/routes"
	serv "github.com/sql-queries/server"
)

func main() {
	serv.Conn().AutoMigrate(&models.User{})
	serv.Conn().AutoMigrate(&models.Session{})
	serv.Conn().AutoMigrate(&models.RealEstate{})
	setupRoutes.Routes()
	ops.Execute()
}
