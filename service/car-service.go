package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchCarData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type getCarDataServie struct{}

func NewCarService() CarService {
	return &getCarDataServie{}
}

func (*getCarDataServie) FetchCarData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)
	//TODO:call external api
	data, _ := client.Get(carServiceUrl)

	//TODO: write respone to channel
	carDataChannel <- data

}
