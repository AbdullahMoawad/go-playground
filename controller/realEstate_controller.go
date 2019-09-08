package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"net/http"
)

func CreateEstate(w http.ResponseWriter, r *http.Request) {
	newRealEstate := *models.NewRealEstate()
	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		fmt.Println(models.Logger(404, "Error getting data from request .. "), err)
	}

	newRealEstate.RealEstateId = uuid.New()

	if err := serv.Conn().Create(&newRealEstate); err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	_ = json.NewEncoder(w).Encode(&newRealEstate)
}

func UpdateEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var realEstate *models.RealEstate

	params := mux.Vars(r)
	estateId := params["estateId"]

	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		panic(err)
	}
	// @todo fix user id change
	if err := serv.Conn().Model(&realEstate).Where("real_estate_id = ?", estateId).Updates(map[string]interface{}{
		"realEstateName":        realEstate.RealEstateName,
		"realEstateType":        realEstate.RealEstateType,
		"categoryName":          realEstate.CategoryName,
		"categoryId":            realEstate.CategoryId,
		"paymentAmount":         realEstate.PaymentAmount,
		"city":                  realEstate.City,
		"floorSpace":            realEstate.FloorSpace,
		"numberOfBalconies":     realEstate.NumberOfBalconies,
		"numberOfBedrooms":      realEstate.NumberOfBedrooms,
		"numberOfBathrooms":     realEstate.NumberOfBathrooms,
		"numberOfGarages":       realEstate.NumberOfGarages,
		"numberOfParkingSpaces": realEstate.NumberOfParkingSpaces,
		"elevator":              realEstate.Elevator,
		"petsAllowed":           realEstate.PetsAllowed,
		"estateDiscribtion":     realEstate.EstateDiscribtion,
		"estatesStatus":         realEstate.EstatesStatus,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}