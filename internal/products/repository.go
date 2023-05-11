package products

import "errors"

var (
	ErrRepoInternal = errors.New("internal error")
	ErrRepoNotFound = errors.New("products not found")
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	var prodListResp []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	for _, product := range prodList {
		if product.SellerID == sellerID {
			prodListResp = append(prodListResp, product)
		}
	}
	if prodListResp == nil {
		return prodListResp, ErrRepoNotFound
	}
	return prodListResp, nil
}
