package service

import (
	"fmt"
	"net/http"
)

type CarOwnerService interface {
	FetchOwnerData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type getCarOwnerDataServie struct{}

func NewCarOwnerService() CarOwnerService {
	return &getCarOwnerDataServie{}
}

func (*getCarOwnerDataServie) FetchOwnerData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)
	//TODO:call external api
	data, _ := client.Get(ownerServiceUrl)

	//TODO: write respone to channel
	ownerDataChannel <- data
}
