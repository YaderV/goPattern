package adapter

// Payment ...
type Payment interface {
	Pay() error
}

// ProcessPayment ...
func ProcessPayment(p Payment) error {
	return p.Pay()
}

// CashPayment ...
type CashPayment struct{}

// Pay ...
func (cp *CashPayment) Pay() error {
	// Do something
	return nil
}

// BankPayment ...
type BankPayment struct{}

// Pay ...
func (bp *BankPayment) Pay(account int) error {
	// Do something
	return nil
}

// BankPaymentAdapter ...
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	Account     int
}

// Pay ...
func (cpa *BankPaymentAdapter) Pay() error {
	return cpa.BankPayment.Pay(cpa.Account)
}
