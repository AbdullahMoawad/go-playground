package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type RealEstate struct {
	gorm.Model
	UserId                uint      `gorm:"primary_key"`
	RealEstateId          uuid.UUID `gorm:"primary_key"`
	RealEstateType        string
	RealEstateName        string
	CategoryName          string
	CategoryId            int
	PaymentAmount         int
	City                  string
	FloorSpace            int
	NumberOfBalconies     int
	NumberOfBedrooms      int
	NumberOfBathrooms     int
	NumberOfGarages       int
	NumberOfParkingSpaces int
	Elevator              string
	//ContractDetails		  *Contract `gorm:"embedded"`
	PetsAllowed       bool
	EstateDiscribtion string
	EstatesStatus     bool
	IsActive          bool
}

type Contract struct {
	Owner           string
	DateOfSigniture string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
