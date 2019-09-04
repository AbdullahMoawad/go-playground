package controller

import (
	"encoding/json"
	"fmt"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"google.golang.org/appengine/user"
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

func UpdateRealEstate(w http.ResponseWriter, r *http.Request) {
	var realEstate *models.RealEstate
	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		panic(err)
	}

	if err := serv.Conn().Model(&realEstate).Where("email = ?", realEstate.Email).Updates(map[string]interface{}{
		"userid":
		"realestateId":
		"categoryName":
		"categoryId":
		"realEstateName":
		"paymentAmount":
		"city":                  string
		"realEstateType":            string
		FloorSpace            int
		NumberOfBalconies     int
		NumberOfBedrooms      int
		NumberOfBathrooms     int
		NumberOfGarages       int
		NumberOfParkingSpaces int
		Elevator			  string
		//ContractDetails		  *Contract `gorm:"embedded"`
		PetsAllowed           bool
		EstateDiscribtion     string
		EstatesStatus         bool
		IsActive              bool}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}