package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/models"
	serv "real-estate/server"
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

	if err := serv.Conn().Model(&realEstate).Where("real_estate_id = ?", estateId).Updates(map[string]interface{}{
		"userId":                realEstate.UserId,
		"realEstateName":        realEstate.Name,
		"realEstateType":        realEstate.Type,
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
		"estateDiscribtion":     realEstate.EstateDescription,
		"estatesStatus":         realEstate.EstatesStatus,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}
