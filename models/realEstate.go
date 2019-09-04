package models

import (
	"github.com/jinzhu/gorm"
)

type RealEstate struct {
	gorm.Model
	UserId                uint `gorm:"primary_key"`
	RealEstateId          uint `gorm:"primary_key"`
	CategoryName 		  string
	CategoryId			  int
	RealEstateName            string
	PaymentAmount         int
	City                  string
	EstateType            string
	FloorSpace            int
	NumberOfBalconies     int
	NumberOfBedrooms      int
	NumberOfBathrooms     int
	NumberOfGarages       int
	NumberOfParkingSpaces int
	Elevator			  string
	//ContractDetails		  *Contract `gorm:"embedded"`
	PetsAllowed           bool
	EstateDiscribtion     string
	EstatesStatus         bool
	IsActive              bool
}

type Contract struct {
	Owner 			string
	DateOfSigniture string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
