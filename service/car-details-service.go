package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danish45007/go-rest/entity"
)

type CarDetailsService interface {
	GetCarDetails() entity.CarDetails
}

var (
	carService       CarService      = NewCarService()
	carOwnerService  CarOwnerService = NewCarOwnerService()
	carDataChannel                   = make(chan *http.Response)
	ownerDataChannel                 = make(chan *http.Response)
)

type carServices struct{}

func NewCarDetailsController() CarDetailsService {
	return &carServices{}
}

func (*carServices) GetCarDetails() entity.CarDetails {
	go carService.FetchCarData()

	go carOwnerService.FetchOwnerData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r2 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r2.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
