package app

import (
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
	"github.com/Khanabeev/banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{service: mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestCustomerHandler(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	t.Run("first", func(t *testing.T) {
		// Arrange
		dummyCustomers := []dto.CustomerResponse{
			{"1001", "Ashish", "New Delhi", "12123012", "2000-01-01", "1"},
			{"1002", "John", "Moscow", "56734", "2001-01-01", "1"},
		}
		mockService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)
		req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

		// Act
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assert
		if recorder.Code != http.StatusOK {
			t.Error("Failed while testing the status code")
		}
	})

	t.Run("get Error", func(t *testing.T) {
		// Arrange
		mockService.EXPECT().GetAllCustomers("").Return(nil, errors2.UnexpectedError("some data"))
		req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

		// Act
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assert
		if recorder.Code != http.StatusInternalServerError {
			t.Error("Failed while testing the status code")
		}
	})

}
