package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
)

type CategoryController struct{}

func (self *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//@todo use requestObject => inside it have validate, format, execute methods
	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		App.JsonLogger(w, 500, common.ErrorMessageFailedToDecodeListRequest, err)
		App.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
		return
	}

	IsExist := category.IsCategoryExist(category.Name)
	if IsExist {
		App.JsonLogger(w, common.StatusNotFound, common.CategoryAlreadyExist, IsExist)
		App.Logger(common.CategoryAlreadyExist, "")
		return
	}
	if err := serv.Conn().Create(&category); err.Error != nil {
		App.JsonLogger(w, 500, "Error create category", err)
		App.Logger("Error create category", "error")
		return
	}
	App.Json(w, category, common.StatusOK)
}

func (self *CategoryController) ListCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var category []models.Category
	queryResult := serv.Conn().Find(&category)
	if queryResult.Error != nil {
		App.JsonLogger(w, 500, "No category found ..", queryResult)
		App.Logger("No category found ..", "error")
		return
	}
	App.Json(w, queryResult, common.StatusOK)

}

func (self *CategoryController) OneCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := common.GetId(r)
	var category []models.Category
	queryResult := serv.Conn().Where("id = ?", id).First(&category)
	if queryResult.Error != nil {
		App.JsonLogger(w, 500, "No category found ..", queryResult)
		App.Logger("No category found ..", "error")
		return
	}
	App.Json(w, queryResult, common.StatusOK)
}

func (self *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	categoryId := common.GetId(r)
	var category []models.Category
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", categoryId).Unscoped().Delete(&category)
	if queryResult.Error != nil {
		App.JsonLogger(w, 500, "No category found ..", queryResult)
		App.Logger("No category found ..", "error")
		return
	}
	App.Json(w, queryResult, common.StatusOK)
}
