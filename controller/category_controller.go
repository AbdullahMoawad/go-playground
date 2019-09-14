package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := serv.Conn().Create(&category); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(&category)
}

func ListCategories(w http.ResponseWriter, r *http.Request) {
	sessionId := common.GetSessionId(r)
	var estates []models.RealEstate

	err, userId := services.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "Error Finding user", err))
		return
	}

	queryResult := serv.Conn().Where("user_id = ?", userId).Find(&estates)
	if queryResult != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func OneCategory(w http.ResponseWriter, r *http.Request) {
	estateId := common.GetId(r)
	var estates []models.RealEstate
	queryResult := serv.Conn().Where("id = ?", estateId).First(&estates)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := common.GetId(r)
	var category []models.Category
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", categoryId).Unscoped().Delete(&category)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}
