package dto

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestValidation(t *testing.T) {
	// AAA principle

	t.Run("Invalid type transaction", func(t *testing.T) {
		// Arrange
		request := NewTransactionRequest{
			TransactionType: "invalid transaction type",
		}
		// Act
		appError := request.Validate()

		// Assert
		require.NotNil(t, appError)
		require.Equal(t, http.StatusUnprocessableEntity, appError.Code)
	})

	t.Run("Invalid amount of money", func(t *testing.T) {
		// Arrange
		request := NewTransactionRequest{
			TransactionType: DEPOSIT,
			Amount:          -100,
		}
		// Act
		appError := request.Validate()

		// Assert
		require.Equal(t, "Amount cannot be negative", appError.Message)
		require.Equal(t, http.StatusUnprocessableEntity, appError.Code)
	})

}
