package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"real-estate/models"
	serv "real-estate/server"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.NewCategory()
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	category.Id = uuid.New()
	if err := serv.Conn().Create(&category); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(&category)
}
