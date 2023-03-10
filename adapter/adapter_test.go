package adapter_test

import (
	"testing"

	"github.com/YaderV/goPattern/adapter"
)

func TestAdapter(t *testing.T) {
	cashPayment := &adapter.CashPayment{}

	err := adapter.ProcessPayment(cashPayment)
	if err != nil {
		t.Fatalf("there should not be an error")
	}

	bankPaymentAdapter := &adapter.BankPaymentAdapter{
		BankPayment: &adapter.BankPayment{},
		Account:     1000,
	}

	err = adapter.ProcessPayment(bankPaymentAdapter)
	if err != nil {
		t.Fatalf("there should not be an error")
	}
}
