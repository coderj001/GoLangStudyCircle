package account

import (
	"testing"
)

func TestDebitCardAccount_Withdraw(t *testing.T) {
	dca := DebitCardAccount{balance: 100}

	dca.Withdraw(50)

	if dca.balance != 50 {
		t.Errorf("Expected balance to be 50, got %v", dca.balance)
	}
}

func TestDebitCardAccount_CanWithdraw(t *testing.T) {
	dca := DebitCardAccount{balance: 100}

	if !dca.CanWithdraw(50) {
		t.Error("Expected CanWithdraw to return true")
	}

	if dca.CanWithdraw(200) {
		t.Error("Expected CanWithdraw to return false")
	}
}

func TestDebitCardAccount_GetIdentifier(t *testing.T) {
	dca := DebitCardAccount{}

	if dca.GetIdentifier() != "DEBIT_CARD_ACCOUNT" {
		t.Errorf("Expected GetIdentifier to return 'DEBIT_CARD_ACCOUNT', got %v", dca.GetIdentifier())
	}
}

func TestCreditCardAccount_Withdraw(t *testing.T) {
	cca := CreditCardAccount{balance: 100, limit: 200}
	cca.Withdraw(50)

	if cca.balance != 50 {
		t.Errorf("Expected balance to be 50, got %v", cca.balance)
	}

}

func TestCreditCardAccount_CanWithdraw(t *testing.T) {
	cca := CreditCardAccount{balance: 100, limit: 200}

	if !cca.CanWithdraw(50) {
		t.Error("Expected CanWithdraw to return true")
	}

	if !cca.CanWithdraw(300) {
		t.Error("Expected CanWithdraw to return true")
	}

	if cca.CanWithdraw(400) {
		t.Error("Expected CanWithdraw to return false")
	}

}

func TestCreditCardAccount_GetIdentifier(t *testing.T) {
	cca := CreditCardAccount{}

	if cca.GetIdentifier() != "CREDIT_CARD_ACCOUNT" {
		t.Errorf("Expected GetIdentifier to return 'CREDIT_CARD_ACCOUNT', got %v", cca.GetIdentifier())
	}

}

func TestPaytmWallet_Withdraw(t *testing.T) {
	pw := PaytmWallet{balance: 100}

	pw.Withdraw(50)

	if pw.balance != 50 {
		t.Errorf("Expected balance to be 50, got %v", pw.balance)
	}

}

func TestPaytmWallet_CanWithdraw(t *testing.T) {
	pw := PaytmWallet{balance: 100}

	if !pw.CanWithdraw(50) {
		t.Error("Expected CanWithdraw to return true")
	}

	if pw.CanWithdraw(200) {
		t.Error("Expected CanWithdraw to return false")
	}

}

func TestPaytmWallet_GetIdentifier(t *testing.T) {
	pw := PaytmWallet{}
	if pw.GetIdentifier() != "PAYTM_WALLET" {
		t.Errorf("Expected GetIdentifier to return 'PAYTM_WALLET', got %v", pw.GetIdentifier())
	}

}
