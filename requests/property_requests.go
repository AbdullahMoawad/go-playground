package requests

import "real-estate/models"

type PropertyRequest struct {
	models.Property
}

func NewPropertyRequest() *PropertyRequest {
	return &PropertyRequest{}
}
