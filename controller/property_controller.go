package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"property/App"
	"property/common"
	"property/common/helpers"
	"property/models"
	"property/server"
)

type PropertyController struct {
	App.Controller
}

func (self PropertyController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	newProperty := models.Property{}

	if err := json.NewDecoder(r.Body).Decode(&newProperty); err != nil {
		self.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest)
		self.Logger(common.ErrorMessageFailedToDecodeListRequest, "error", err)
		return
	}

	sessionId := models.GetCurrentSessionId(r)

	err, userId := helpers.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		self.JsonLogger(w, 404, err.Error())
		self.Logger("error", err.Error(), err)
		return
	}
	newProperty.UserId = userId
	newProperty.Id = uuid.New()

	if err := server.CreatePostgresDbConnection().Create(&newProperty); err.Error != nil {
		return
	}
	self.Json(w, &newProperty, common.StatusOK)
}

func (self PropertyController) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var property *models.Property
	userId := helpers.GetCurrentUserId(r)

	if err := json.NewDecoder(r.Body).Decode(&property); err != nil {
		self.JsonLogger(w, common.StatusBadRequest, common.ErrorMessageFailedToDecodeListRequest)
		self.Logger(common.ErrorMessageFailedToDecodeListRequest, "error", err)
		return
	}

	if err := server.CreatePostgresDbConnection().Model(&property).Where("id = ?", userId).Updates(map[string]interface{}{
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
	self.Json(w, "updated successfully", common.StatusOK)

}

func (self PropertyController) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	sessionId := models.GetCurrentSessionId(r)

	err, userId := helpers.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		self.JsonLogger(w, common.StatusNotFound, "Error Finding user id")
		self.Logger("Error Finding user", "error", err)
		return
	}

	queryResult := models.FindAllProperties(userId)
	if queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property fount ..")
		self.Logger("Error finding user", "error", queryResult.Error.Error())
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self PropertyController) One(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	propertyId := helpers.GetCurrentUserId(r)
	var property []models.Property

	queryResult := server.CreatePostgresDbConnection().Where("id = ?", propertyId).First(&property)
	if queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property found ..")
		self.Logger("Error finding property", "error", queryResult.Error.Error())
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self PropertyController) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	propertyId := helpers.GetCurrentPropertyId(r)
	var property []models.Property

	// unscoped to permanently delete record from database
	if queryResult := server.CreatePostgresDbConnection().Where("id = ?", propertyId).Unscoped().Delete(&property); queryResult.Error != nil {
		self.JsonLogger(w, 404, "No property found ..")
		self.Logger("Error finding property", "error", queryResult.Error.Error())
		return
	}
	self.Json(w, "deleted successfully", common.StatusOK)
}
