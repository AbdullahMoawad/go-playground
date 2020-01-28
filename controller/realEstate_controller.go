package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
)

type RealEstateController struct{}

func (self RealEstateController) CreateRealEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newRealEstate := models.NewRealEstate()

	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		App.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest, err)
		App.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
		return
	}

	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		App.JsonLogger(w, 404, err.Error(), err)
		App.Logger(err.Error(), "error")
		return
	}
	newRealEstate.UserId = userId
	newRealEstate.Id = uuid.New()

	if err := serv.Conn().Create(&newRealEstate); err.Error != nil {
		return
	}
	json.NewEncoder(w).Encode(&newRealEstate)
}

func (self RealEstateController) UpdateRealEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var realEstate *models.RealEstate
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		App.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest, err)
		App.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
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

func (self RealEstateController) ListRealEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		App.JsonLogger(w, common.StatusNotFound, "Error Finding user id", err)
		App.Logger("Error Finding user", "error")
		return
	}

	queryResult := services.FindAllEstates(userId)
	if queryResult.Error != nil {
		App.JsonLogger(w, 404, "No real estates fount ..", queryResult)
		App.Logger("Error finding user", "error")
		return
	}

	json.NewEncoder(w).Encode(queryResult)
}

func (self RealEstateController) OneRealEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	estateId := common.GetId(r)
	var estates []models.RealEstate
	queryResult := serv.Conn().Where("id = ?", estateId).First(&estates)
	if queryResult.Error != nil {
		App.JsonLogger(w, 404, "No real estates found ..", queryResult)
		App.Logger("Error finding realestate", "error")
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func (self RealEstateController) DeleteRealEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	estateId := common.GetId(r)
	var estates []models.RealEstate
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", estateId).Unscoped().Delete(&estates)
	if queryResult.Error != nil {
		App.JsonLogger(w, 404, "No real estates found ..", queryResult)
		App.Logger("Error finding realestate", "error")
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}
