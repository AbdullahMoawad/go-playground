package models

import (
	"github.com/google/uuid"
	"time"
)

type RealEstate struct {
	UserId                uint      `json:"userId"`
	Id                    uuid.UUID `gorm:"primary_key" json:"id"`
	Type                  string    `json:"type"`
	Name                  string    `json:"name"`
	CategoryName          string    `json:"categoryName"`
	CategoryId            int       `json:"categoryId"`
	PaymentAmount         int       `json:"paymentAmount"`
	City                  string    `json:"city"`
	FloorSpace            int       `json:"floorSpace"`
	NumberOfBalconies     int       `json:"balconies"`
	NumberOfBedrooms      int       `json:"bedrooms"`
	NumberOfBathrooms     int       `json:"bathrooms"`
	NumberOfGarages       int       `json:"garages"`
	NumberOfParkingSpaces int       `json:"parkingSpaces"`
	Elevator              string    `json:"elevator"`
	PetsAllowed           bool      `json:"petsAllowed"`
	EstateDescription     string    `json:"description"`
	EstatesStatus         bool      `json:"status"`
	IsActive              bool      `json:"isActive"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time `sql:"index"`
}

type Contract struct {
	Owner           string
	DateOfSigniture string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
