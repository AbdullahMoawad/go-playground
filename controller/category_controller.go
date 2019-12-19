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
	defer r.Body.Close()

	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		json.NewEncoder(w).Encode(models.Logger(500, common.DecodingError, err))
		return
	}
	IsExist := services.IsExist(category.Name)
	if IsExist {
		json.NewEncoder(w).Encode(models.Logger(500, common.CategoryAlreadyExist, nil))
		return
	}
	if err := serv.Conn().Create(&category); err.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(500, "Error create category", err.Error))
		return
	}
	json.NewEncoder(w).Encode(&category)
}

func ListCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var category []models.Category
	queryResult := serv.Conn().Find(&category)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func OneCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := common.GetId(r)
	var category []models.Category
	queryResult := serv.Conn().Where("id = ?", id).First(&category)
	if queryResult.Error != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "No real estates fount ..", queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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
