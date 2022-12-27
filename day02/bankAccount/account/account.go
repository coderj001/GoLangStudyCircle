package account

// TDD Bank Account app

type Account interface {
	Withdraw(amount float64)
	CanWithdraw(amount float64) bool
	GetIdentifier() string
}
type DebitCardAccount struct {
	balance float64
}

func (dca *DebitCardAccount) Withdraw(amount float64) {
	dca.balance -= amount
}

func (dca *DebitCardAccount) CanWithdraw(amount float64) bool {
	return dca.balance >= amount
}

func (dca *DebitCardAccount) GetIdentifier() string {
	return "DEBIT_CARD_ACCOUNT"
}

type CreditCardAccount struct {
	balance float64
	limit   float64
}

func (cca *CreditCardAccount) Withdraw(amount float64) {
	cca.balance -= amount
}

func (cca *CreditCardAccount) CanWithdraw(amount float64) bool {
	return cca.balance+cca.limit >= amount
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}

type PaytmWallet struct {
	balance float64
}

func (pw *PaytmWallet) Withdraw(amount float64) {
	pw.balance -= amount
}

func (pw *PaytmWallet) CanWithdraw(amount float64) bool {
	return pw.balance >= amount
}

func (pw *PaytmWallet) GetIdentifier() string {
	return "PAYTM_WALLET"
}
