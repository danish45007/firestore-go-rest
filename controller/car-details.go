package controller

import (
	"encoding/json"
	"net/http"

	"github.com/danish45007/go-rest/service"
)

type controller struct{}

var (
	carDetailsService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetCarDetails()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
