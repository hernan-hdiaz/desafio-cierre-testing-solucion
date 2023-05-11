package products

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandlerGetByID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Arrange.
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `[
			{
				"id": "mock",
				"seller_id": "FEX112AC",
				"description": "generic product",
				"price": 123.55
			}
		]
		`
		var prodListExpected []Product
		prodListExpected = append(prodListExpected, Product{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		})

		service := NewServiceMockTestify()
		service.On("GetAllBySeller", "bla").Return(prodListExpected, nil)

		handler := NewHandler(service)

		//router := gin.New()
		router := gin.Default()

		router.GET("/api/v1/products", handler.GetProducts)

		responseRecorder := httptest.NewRecorder()

		// Act.
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?seller_id=bla",
			nil,
		))

		// Assert.
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
		service.AssertExpectations(t)
	})

	t.Run("seller_id required", func(t *testing.T) {
		// Arrange.
		expectedHTTPStatusCode := http.StatusBadRequest
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		expectedResponse := "{\"error\":\"seller_id query param is required\"}"

		service := NewServiceMockTestify()
		handler := NewHandler(service)

		//router := gin.New()
		router := gin.Default()

		router.GET("/api/v1/products", handler.GetProducts)

		responseRecorder := httptest.NewRecorder()
		//err := errors.New("seller_id query param is required")
		// Act.
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?seller_id=",
			nil,
		))

		// Assert.
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.Equal(t, expectedResponse, responseRecorder.Body.String())
	})
	t.Run("product not found", func(t *testing.T) {
		// Arrange.
		expectedHTTPStatusCode := http.StatusInternalServerError
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		service := NewServiceMockTestify()
		service.On("GetAllBySeller", "bla").Return([]Product{}, ErrServiceNotFound)

		handler := NewHandler(service)

		router := gin.New()
		//router := gin.Default()

		router.GET("/api/v1/products", handler.GetProducts)
		responseRecorder := httptest.NewRecorder()
		expectedResponse := "{\"error\":\"products not found\"}"

		// Act.
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?seller_id=bla",
			nil,
		))

		// Assert.
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.Equal(t, expectedResponse, responseRecorder.Body.String())
		service.AssertExpectations(t)
	})
}
