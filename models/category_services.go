package models

import (
	"real-estate/server"
)

func (self *Category) FindCategoryByName(name string) *Category {
	newCategory := &Category{}
	queryResult := server.Conn().Where(&Category{Name: name}).Find(&newCategory)
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
