package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAllBySeller_Testify(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		var prodListExpected []Product
		prodListExpected = append(prodListExpected, Product{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		})

		rp := NewRepositoryMockTestify()
		rp.
			On("GetAllBySeller", "FEX112AC").
			Return(prodListExpected, nil)

		sv := NewService(rp)

		// act
		prodList, err := sv.GetAllBySeller("FEX112AC")

		// assert
		assert.NoError(t, err)
		assert.Equal(t, prodListExpected, prodList)
		// -> expected calls
		rp.AssertExpectations(t)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		rp := NewRepositoryMockTestify()
		rp.
			On("GetAllBySeller", "any").
			Return([]Product{}, ErrRepoNotFound)

		sv := NewService(rp)

		// act
		prodList, err := sv.GetAllBySeller("any")

		// assert
		assert.Empty(t, prodList)
		assert.ErrorIs(t, err, ErrServiceNotFound)
		// -> expected calls
		rp.AssertExpectations(t)
	})

}
