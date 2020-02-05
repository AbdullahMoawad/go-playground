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

	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		self.JsonLogger(w, 500, common.ErrorMessageFailedToDecodeListRequest)
		self.Logger("error", common.ErrorMessageFailedToDecodeListRequest, err)
		return
	}

	IsExist := category.IsCategoryExist(category.Name)
	if IsExist {
		self.JsonLogger(w, common.StatusNotFound, common.CategoryAlreadyExist)
		self.Logger("error", common.CategoryAlreadyExist, nil)
		return
	}
	if queryResult := serv.CreatePostgresDbConnection().Create(&category); queryResult.Error != nil {
		self.JsonLogger(w, 500, "Error create category")
		self.Logger("error", "Error create category", queryResult.Error.Error())
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

	queryResult := serv.CreatePostgresDbConnection().Find(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..")
		self.Logger("error", "No category found ..", queryResult.Error.Error())
		return
	}
	self.Json(w, queryResult, common.StatusOK)

}

func (self *CategoryController) One(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var category []models.Category

	categoryId := models.GetCurrentCategoryId(r)

	queryResult := serv.CreatePostgresDbConnection().Where("id = ?", categoryId).First(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..")
		self.Logger("error", "No category found ..", queryResult.Error.Error())
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}

func (self *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	categoryId := models.GetCurrentCategoryId(r)
	var category []models.Category

	// unscoped to permanently delete record from database
	queryResult := serv.CreatePostgresDbConnection().Where("id = ?", categoryId).Unscoped().Delete(&category)
	if queryResult.Error != nil {
		self.JsonLogger(w, 500, "No category found ..")
		self.Logger("error", "No category found ..", queryResult.Error.Error())
		return
	}
	self.Json(w, queryResult, common.StatusOK)
}
