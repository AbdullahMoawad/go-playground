package controller

import (
	"encoding/json"
	"net/http"
	"real-estate/App"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
)

type CategoryController struct {
	App.Controller
}

func (self *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//@todo use requestObject => inside it have validate, format, execute methods
	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		self.JsonLogger(w, 500, common.ErrorMessageFailedToDecodeListRequest, err)
		self.Logger(common.ErrorMessageFailedToDecodeListRequest, "error")
		return
	}

	IsExist := category.IsCategoryExist(category.Name)
	if IsExist {
		self.JsonLogger(w, common.StatusNotFound, common.CategoryAlreadyExist, IsExist)
		self.Logger(common.CategoryAlreadyExist, "")
		return
	}
	if err := serv.Conn().Create(&category); err.Error != nil {
		self.JsonLogger(w, 500, "Error create category", err)
		self.Logger("Error create category", "error")
		return
	}
	self.Json(w, category, common.StatusOK)
}

func (self *CategoryController) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var category []models.Category

	// if want to show a specific number of categories
	// queryResult := serv.Conn().Limit(4).Find(&category)

	// if want to print the number of  categories
	// var count init64
	// queryResult := serv.Conn().Find(&category).count(&count)

	queryResult := serv.Conn().Find(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..", queryResult)
		self.Logger("No category found ..", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)

}

func (self *CategoryController) One(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := common.GetId(r)
	var category []models.Category
	queryResult := serv.Conn().Where("id = ?", id).First(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..", queryResult)
		self.Logger("No category found ..", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	categoryId := common.GetId(r)
	var category []models.Category
	// unscoped to permanently delete record from database
	queryResult := serv.Conn().Where("id = ?", categoryId).Unscoped().Delete(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..", queryResult)
		self.Logger("No category found ..", "error")
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}
