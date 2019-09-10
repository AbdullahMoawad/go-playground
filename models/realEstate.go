package models

import (
	"github.com/google/uuid"
	"time"
)

type RealEstate struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	Id            uuid.UUID  `gorm:"primary_key" json:"id"`
	UserId        uint       `gorm:"primary_key" json:"userId"`
	Type          string     `json:"Type"`
	Name          string     `json:"name"`
	CategoryName  string     `json:"categoryName"`
	CategoryId    int        `json:"categoryId"`
	PaymentAmount int        `json:"paymentAmount"`
	City          string     `json:"city"`
	FloorSpace    int        `json:"floorSpace"`
	Balconies     int        `json:"balconies"`
	Bedrooms      int        `json:"bedrooms"`
	Bathrooms     int        `json:"bathrooms"`
	Garages       int        `json:"garages"`
	ParkingSpaces int        `json:"parkingSpaces"`
	Elevator      string     `json:"elevator"`
	PetsAllowed   bool       `json:"petsAllowed"`
	Description   string     `json:"description"`
	Status        bool       `json:"status"`
	IsActive      bool       `json:"isActive"`
}

type Contract struct {
	Owner           string
	DateOfSigniture string
}

func NewRealEstate() *RealEstate {
	var realEstate RealEstate
	return &realEstate
}
