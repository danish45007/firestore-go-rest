package service

import "github.com/danish45007/go-rest/entity"

type carDetailsService interface {
	GetCarDetails() entity.CarDetails
}

type service struct{}

func NewCarDetailsController() carDetailsService {
	return &service{}
}

func (*service) GetCarDetails() entity.CarDetails {

}
