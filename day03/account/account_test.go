package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestInsufficientBalanceError(t *testing.T) {
	acc := Account{balance: 50}
	err := acc.Withdraw(100)
	if err == nil {
		t.Error("Expected error to be returned")
	}
	expectedError := "insufficient balance, current balance: 50.00, amount needed: 100.00"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %s, got %s", expectedError, err.Error())
	}
}
