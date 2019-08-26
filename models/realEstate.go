package models

import (
	"github.com/jinzhu/gorm"
)

type RealEstate struct {
	gorm.Model
	UserId                uint `gorm:"primary_key"`
	EstateName            string
	PaymentAmount			int
	City                  string
	EstateType            string
	FloorSpace				int
	NumberOfBalconies     int
	NumberOfBedrooms      int
	NumberOfBathrooms     int
	NumberOfGarages       int
	NumberOfParkingSpaces int
	PetsAllowed           bool
	EstateDiscribtion     string
	EstatesStatus         bool
	IsActive 			  bool
}

func NewRealEstate() *RealEstate  {
	var realEstate RealEstate
	return &realEstate
}