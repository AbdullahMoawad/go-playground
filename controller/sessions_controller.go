package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
)

func CreateSession (userId uint){
	session := models.Session{}
	session.UserId = userId
	session.SessionId = uuid.New()
	fmt.Println(session)
	if err := serv.Conn().Create(&session); err != nil {
		fmt.Println("error while creating session")
		return
	}
}