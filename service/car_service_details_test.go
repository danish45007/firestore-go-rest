package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	carDetailServices = NewCarDetailsController()
)

func TestGetDetails(t *testing.T) {
	carDetails := carDetailServices.GetCarDetails()
	fmt.Println(carDetails)
	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.ID)
	assert.Equal(t, "Mitsubishi", carDetails.Brand)
}
