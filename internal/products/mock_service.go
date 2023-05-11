package products

import (
	"github.com/stretchr/testify/mock"
)

// constructor
func NewServiceMockTestify() *ServiceMockTestify {
	return &ServiceMockTestify{}
}

// RepositoryMockTestify is a product repository mock.
type ServiceMockTestify struct {
	mock.Mock
}

func (sv *ServiceMockTestify) GetAllBySeller(sellerID string) ([]Product, error) {
	// input
	// -> for expectations (to check if the input is the same as the expected)
	args := sv.Called(sellerID)

	// output
	productList := args.Get(0).([]Product)
	err := args.Error(1)
	return productList, err
}
