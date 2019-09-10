package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sql-queries/models"
	serv "github.com/sql-queries/server"
	"net/http"
)

func CreateEstate(w http.ResponseWriter, r *http.Request) {
	newRealEstate := models.NewRealEstate()

	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	newRealEstate.Id = uuid.New()
	if err := serv.Conn().Create(&newRealEstate); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(&newRealEstate)
}

func UpdateEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var realEstate *models.RealEstate
	params := mux.Vars(r)
	estateId := params["estateId"]

	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		panic(err)
	}

	if err := serv.Conn().Model(&realEstate).Where("id = ?", estateId).Updates(map[string]interface{}{
		"userId":                realEstate.UserId,
		"name":                  realEstate.Name,
		"type":                  realEstate.Type,
		"categoryName":          realEstate.CategoryName,
		"categoryId":            realEstate.CategoryId,
		"paymentAmount":         realEstate.PaymentAmount,
		"city":                  realEstate.City,
		"floorSpace":            realEstate.FloorSpace,
		"balconies":             realEstate.Balconies,
		"bedrooms":              realEstate.Bedrooms,
		"bathrooms":             realEstate.Bathrooms,
		"garages":               realEstate.Garages,
		"numberOfParkingSpaces": realEstate.ParkingSpaces,
		"elevator":              realEstate.Elevator,
		"petsAllowed":           realEstate.PetsAllowed,
		"description":           realEstate.Description,
		"status":                realEstate.Status,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}
