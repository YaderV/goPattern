package factory

import (
	"fmt"
)

const (
	// InvoiceStatusPending ...
	InvoiceStatusPending int = iota
	// InvoiceStatusPaid ...
	InvoiceStatusPaid
)

const (
	// PaymentMethodCash ...
	PaymentMethodCash int = iota
	// PaymentMethodCC ...
	PaymentMethodCC
)

// Invoice struct
type Invoice struct {
	ID            int
	Amount        int
	Status        int
	PaymentMethod int
}

// IPayment ...
type IPayment interface {
	Pay() error
}

// PaymentBase ...
type PaymentBase struct {
	invoice *Invoice
}

// Validate ..
func (pm *PaymentBase) Validate() error {
	if pm.invoice.Status == PaymentMethodCash {
		return fmt.Errorf("this invoice is already paid")
	}
	return nil
}

// PMCash ...
type PMCash struct {
	PaymentBase
}

// Pay ...
func (pc *PMCash) Pay() error {
	err := pc.Validate()
	if err != nil {
		return err
	}
	pc.invoice.Status = InvoiceStatusPaid
	pc.invoice.PaymentMethod = PaymentMethodCash
	// ...
	fmt.Println("We paid with cash")
	return nil
}

// PMCC ...
type PMCC struct {
	PaymentBase
}

// Pay ...
func (pcc *PMCC) Pay() error {
	err := pcc.Validate()
	if err != nil {
		return err
	}
	pcc.invoice.Status = InvoiceStatusPaid
	pcc.invoice.PaymentMethod = PaymentMethodCash
	// ...
	fmt.Println("We paid with credit card")
	return nil
}

// NewPayment ...
func NewPayment(i *Invoice, pm int) (IPayment, error) {
	if pm == PaymentMethodCash {
		return &PMCash{PaymentBase{invoice: i}}, nil
	}

	if pm == PaymentMethodCC {
		return &PMCC{PaymentBase{invoice: i}}, nil
	}

	return nil, fmt.Errorf("wrong payment selected")
}
