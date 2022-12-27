package quickcash

type Withdrawable interface {
	CanWithDraw(amount float64) bool
	WithDraw(amount float64)
	GetIdentifier() string
}

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return sa.balance >= amount
}

func (sa *SavingsAccount) WithDraw(amount float64) {
	sa.balance -= amount
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}

type CreditCardAccount struct {
	limit float64
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return cca.limit >= amount
}

func (cca *CreditCardAccount) WithDraw(amount float64) {
	cca.limit -= amount
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
