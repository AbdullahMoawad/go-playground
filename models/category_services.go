package models

import (
	"github.com/gorilla/mux"
	"net/http"
	"real-estate/server"
)

func (self *Category) FindCategoryByName(name string) *Category {
	newCategory := &Category{}
	queryResult := server.CreatePostgresDbConnection().Where(&Category{Name: name}).Find(&newCategory)
	if queryResult.Error != nil {
		return nil
	}
	return newCategory
}

func (self *Category) IsCategoryExist(name string) bool {
	userRow := self.FindCategoryByName(name)
	if userRow == nil {
		return false
	}
	return true
}

func GetCurrentCategoryId(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}
