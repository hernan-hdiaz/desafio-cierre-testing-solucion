package products

import "github.com/stretchr/testify/mock"

// constructor
func NewRepositoryMockTestify() *RepositoryMockTestify {
	return &RepositoryMockTestify{}
}

// RepositoryMockTestify is a product repository mock.
type RepositoryMockTestify struct {
	mock.Mock
}

func (rp *RepositoryMockTestify) GetAllBySeller(sellerID string) ([]Product, error) {
	// input
	// -> for expectations (to check if the input is the same as the expected)
	args := rp.Called(sellerID)

	// output
	productList := args.Get(0).([]Product)
	err := args.Error(1)
	return productList, err
}
