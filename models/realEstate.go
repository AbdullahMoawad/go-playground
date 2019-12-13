package models

import (
	"github.com/google/uuid"
	"time"
)

type RealEstate struct {
	UserId        string    `json:"userId"`
	Id            uuid.UUID `gorm:"primary_key" json:"id"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	CategoryName  string    `json:"categoryName"`
	CategoryId    int       `json:"categoryId"`
	PaymentAmount int       `json:"paymentAmount"`
	City          string    `json:"city"`
	FloorSpace    int       `json:"floorSpace"`
	Balconies     int       `json:"balconies"`
	Bedrooms      int       `json:"bedrooms"`
	Bathrooms     int       `json:"bathrooms"`
	Garages       int       `json:"garages"`
	ParkingSpaces int       `json:"parkingSpaces"`
	Elevator      string    `json:"elevator"`
	PetsAllowed   bool      `json:"petsAllowed"`
	Description   string    `json:"description"`
	Status        bool      `json:"status"`
	IsActive      bool      `json:"isActive"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

type Contract struct {
	Owner           string
	DateOfSignature string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
