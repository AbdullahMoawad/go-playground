package models

type UserLogin struct {
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
	SessionId string `json:"sessionId"`
	IsActive  bool   `json:"isActive"`
}
