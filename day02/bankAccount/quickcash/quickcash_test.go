package quickcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	fpa := &SavingsAccount{balance: 500}
	fsa := &CreditCardAccount{limit: 500}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {

	fpa := &SavingsAccount{balance: 0}
	fsa := &CreditCardAccount{limit: 500}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}
