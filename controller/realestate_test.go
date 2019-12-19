package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, _ := http.NewRequest("GET", "/estate/all", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListEstates)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Value":[{"userId":"28523200-73a9-4c60-b33a-2772109737b0","id":"55c38ce5-2d4b-4321-a959-592e19f5abaa","type":"22","name":"bork2r","categoryName":"abrag","categoryId":2,"paymentAmount":3,"city":"cairo/egypt","floorSpace":533333,"balconies":5,"bedrooms":3,"bathrooms":2,"garages":1,"parkingSpaces":1,"elevator":"2 elivators","petsAllowed":true,"description":"","status":true,"isActive":false,"CreatedAt":"2019-12-12T21:05:34.927079+02:00","UpdatedAt":"2019-12-12T21:08:27.445022+02:00","DeletedAt":null}],"Error":null,"RowsAffected":1}`
	actual := rr.Body.String()
	assert.Equal(t, expected, actual, "both actual and expected should be the same .")

}

func TestGetEntryByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/estate/one/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "55c38ce5-2d4b-4321-a959-592e19f5abaa")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(OneEstate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
	expected := `{"Value":[{"userId":"28523200-73a9-4c60-b33a-2772109737b0","id":"55c38ce5-2d4b-4321-a959-592e19f5abaa","type":"22","name":"bork2r","categoryName":"abrag","categoryId":2,"paymentAmount":3,"city":"cairo/egypt","floorSpace":533333,"balconies":5,"bedrooms":3,"bathrooms":2,"garages":1,"parkingSpaces":1,"elevator":"2 elivators","petsAllowed":true,"description":"","status":true,"isActive":false,"CreatedAt":"2019-12-12T21:05:34.927079+02:00","UpdatedAt":"2019-12-12T21:08:27.445022+02:00","DeletedAt":null}],"Error":null,"RowsAffected":1}`
	actual := rr.Body.String()
	assert.Equal(t, expected, actual, "both actual and expected should be the same .")
}
