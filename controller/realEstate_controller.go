package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
)

func CreateEstate(w http.ResponseWriter, r *http.Request) {
	newRealEstate := models.NewRealEstate()

	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, common.DecodingError, err))
		return
	}

	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "Error while getting user from header ", err))
		return
	}
	newRealEstate.UserId = userId
	newRealEstate.Id = uuid.New()

	if err := serv.Conn().Create(&newRealEstate); err.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(500, "Error create category", err.Error))
		return
	}
	json.NewEncoder(w).Encode(&newRealEstate)
}

func UpdateEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var realEstate *models.RealEstate
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, common.DecodingError, err))
		return
	}

	if err := serv.Conn().Model(&realEstate).Where("id = ?", id).Updates(map[string]interface{}{
		"name":          realEstate.Name,
		"type":          realEstate.Type,
		"categoryName":  realEstate.CategoryName,
		"categoryId":    realEstate.CategoryId,
		"paymentAmount": realEstate.PaymentAmount,
		"city":          realEstate.City,
		"floorSpace":    realEstate.FloorSpace,
		"balconies":     realEstate.Balconies,
		"bedrooms":      realEstate.Bedrooms,
		"bathrooms":     realEstate.Bathrooms,
		"Garages":       realEstate.Garages,
		"parkingSpaces": realEstate.ParkingSpaces,
		"elevator":      realEstate.Elevator,
		"petsAllowed":   realEstate.PetsAllowed,
		"description":   realEstate.Description,
		"Status":        realEstate.Status,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func ListEstates(w http.ResponseWriter, r *http.Request) {
	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "Error Finding user", err))
		return
	}

	queryResult := services.FindAllEstates(userId)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}

	json.NewEncoder(w).Encode(queryResult)
}

func OneEstate(w http.ResponseWriter, r *http.Request) {
	estateId := common.GetId(r)
	var estates []models.RealEstate
	queryResult := serv.Conn().Where("id = ?", estateId).First(&estates)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func DeleteEstate(w http.ResponseWriter, r *http.Request) {
	estateId := common.GetId(r)
	var estates []models.RealEstate
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", estateId).Unscoped().Delete(&estates)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}
