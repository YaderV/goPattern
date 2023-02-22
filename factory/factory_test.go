package factory_test

import (
	"testing"

	"github.com/YaderV/goPattern/factory"
)

func TestPM(t *testing.T) {
	invoice := &factory.Invoice{
		ID:     1,
		Amount: 1000,
	}

	pm, err := factory.NewPayment(invoice, factory.PaymentMethodCash)
	if err != nil {
		t.Fatal(err)
	}

	err = pm.Pay()
	if err != nil {
		t.Fatal(err)
	}

	if invoice.Status != factory.InvoiceStatusPaid {
		t.Fatalf("status expected: %d; got: %d", factory.InvoiceStatusPaid, invoice.Status)
	}

	if invoice.PaymentMethod != factory.PaymentMethodCash {
		t.Fatalf("method expected: %d; got: %d", factory.PaymentMethodCash, invoice.PaymentMethod)
	}

}
