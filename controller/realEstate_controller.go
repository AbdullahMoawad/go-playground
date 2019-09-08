package controller

import (
	"encoding/json"
	"fmt"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"net/http"
)

func CreateRealEstate(w http.ResponseWriter, r *http.Request) {
	newRealEstate := *models.NewRealEstate()
	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		fmt.Println(models.Logger(404,"Error getting data from request .. " ),err)
	}
	if err := serv.Conn().Create(&newRealEstate); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	_ = json.NewEncoder(w).Encode(&newRealEstate)
}
