package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type RealEstate struct {
	gorm.Model
	UserId                uint      `gorm:"primary_key" json:"userId"`
	Id          uuid.UUID `gorm:"primary_key" json:"realEstateId"`
	Type        string    `json:"Type"`
	Name        string    `json:"realEstateName"`
	CategoryName          string    `json:"categoryName"`
	CategoryId            int       `json:"categoryId"`
	PaymentAmount         int       `json:"paymentAmount"`
	City                  string    `json:"city"`
	FloorSpace            int       `json:"floorSpace"`
	NumberOfBalconies     int       `json:"numberOfBalconies"`
	NumberOfBedrooms      int       `json:"numberOfBedrooms"`
	NumberOfBathrooms     int       `json:"numberOfBathrooms"`
	NumberOfGarages       int       `json:"numberOfGarages"`
	NumberOfParkingSpaces int       `json:"numberOfParkingSpaces"`
	Elevator              string    `json:"elevator"`
	PetsAllowed           bool      `json:"petsAllowed"`
	EstateDescription     string    `json:"estateDescription"`
	EstatesStatus         bool      `json:"estatesStatus"`
	IsActive              bool      `json:"isActive"`
}

type Contract struct {
	Owner           string
	DateOfSigniture string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
