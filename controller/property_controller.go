package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	"real-estate/requests"
	serv "real-estate/server"
	"real-estate/services"
)

type PropertyController struct {
	App.Controller
}

func (self PropertyController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newProperty := requests.NewPropertyRequest()

	if err := json.NewDecoder(r.Body).Decode(&newProperty); err != nil {
		self.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest, err)
		self.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
		return
	}

	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		self.JsonLogger(w, 404, err.Error(), err)
		self.Logger(err.Error(), "error")
		return
	}
	newProperty.UserId = userId
	newProperty.Id = uuid.New()

	if err := serv.Conn().Create(&newProperty); err.Error != nil {
		return
	}
	self.Json(w, &newProperty, common.StatusOK)
}

func (self PropertyController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var property *models.Property
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&property); err != nil {
		self.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest, err)
		self.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
		return
	}

	if err := serv.Conn().Model(&property).Where("id = ?", id).Updates(map[string]interface{}{
		"name":          property.Name,
		"type":          property.Type,
		"categoryName":  property.CategoryName,
		"categoryId":    property.CategoryId,
		"paymentAmount": property.PaymentAmount,
		"city":          property.City,
		"floorSpace":    property.FloorSpace,
		"balconies":     property.Balconies,
		"bedrooms":      property.Bedrooms,
		"bathrooms":     property.Bathrooms,
		"Garages":       property.Garages,
		"parkingSpaces": property.ParkingSpaces,
		"elevator":      property.Elevator,
		"petsAllowed":   property.PetsAllowed,
		"description":   property.Description,
		"Status":        property.Status,
	}); err != nil {
		self.Json(w, &err, common.StatusOK)
		return
	}
}

func (self PropertyController) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessionId := common.GetSessionId(r)

	err, userId := common.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		self.JsonLogger(w, common.StatusNotFound, "Error Finding user id", err)
		self.Logger("Error Finding user", "error")
		return
	}

	queryResult := services.FindAllProperties(userId)
	if queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property fount ..", queryResult)
		self.Logger("Error finding user", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self PropertyController) One(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	propertyId := common.GetId(r)
	var property []models.Property
	queryResult := serv.Conn().Where("id = ?", propertyId).First(&property)
	if queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property found ..", queryResult)
		self.Logger("Error finding property", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self PropertyController) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	propertyId := common.GetId(r)
	var property []models.Property
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", propertyId).Unscoped().Delete(&property)
	if queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property found ..", queryResult)
		self.Logger("Error finding property", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}
