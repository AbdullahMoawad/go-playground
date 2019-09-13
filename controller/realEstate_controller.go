package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"real-estate/common"
	"real-estate/models"
	serv "real-estate/server"
	"real-estate/services"
)

func GetCurrentUserFromHeaders(SessionID string) (models.Log, string) {
	user := &models.User{}
	queryResult := serv.Conn().Where(&models.Session{SessionId: SessionID}).First(user)
	if queryResult.Error != nil {
		return models.Logger(404, "Erro finding user in database", queryResult.Error), ""
	} else {
		return models.Logger(0, "", queryResult.Error), user.Id
	}
}

func CreateEstate(w http.ResponseWriter, r *http.Request) {
	newRealEstate := models.NewRealEstate()

	if err := json.NewDecoder(r.Body).Decode(&newRealEstate); err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, " Error decoding", err))
		fmt.Println()

		return
	}
	sessionId := common.GetSessionId(r)
	err, userId := GetCurrentUserFromHeaders(sessionId)
	if err.Error != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	newRealEstate.UserId = userId

	newRealEstate.Id = uuid.New()
	serv.Conn().Create(&newRealEstate)
	json.NewEncoder(w).Encode(&newRealEstate)
}

func UpdateEstate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var realEstate *models.RealEstate
	id := common.GetId(r)

	if err := json.NewDecoder(r.Body).Decode(&realEstate); err != nil {
		panic(err)
	}

	if err := serv.Conn().Model(&realEstate).Where("id = ?", id).Updates(map[string]interface{}{
		"userId":        realEstate.UserId,
		"name":          realEstate.Name,
		"type":          realEstate.Type,
		"categoryName":  realEstate.CategoryName,
		"categoryId":    realEstate.CategoryId,
		"paymentAmount": realEstate.PaymentAmount,
		"city":          realEstate.City,
		"floorSpace":    realEstate.FloorSpace,
		"balconies":     realEstate.Balconies,
		"bedrooms":      realEstate.Bedrooms,
		"bathrooms":     realEstate.Bathrooms,
		"Garages":       realEstate.Garages,
		"parkingSpaces": realEstate.ParkingSpaces,
		"elevator":      realEstate.Elevator,
		"petsAllowed":   realEstate.PetsAllowed,
		"description":   realEstate.Description,
		"Status":        realEstate.Status,
	}); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}
// @todo list all
func All(w http.ResponseWriter, r *http.Request) {
	sessionId := common.GetSessionId(r)
	var estates []models.RealEstate

	err, userId := services.GetCurrentUserIdFromHeaders(sessionId)
	if err != nil{
		json.NewEncoder(w).Encode(models.Logger(404,"Error Finding user",err))
		return
	}

	queryResult := serv.Conn().Where("user_id = ?", userId).Find(&estates)
	if queryResult != nil{
		json.NewEncoder(w).Encode(models.Logger(404,"No real estates fount ..",queryResult.Error))
		return
	}
	json.NewEncoder(w).Encode(queryResult)
}
