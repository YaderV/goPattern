package observer_test

import (
	"testing"
	"time"

	"github.com/YaderV/goPattern/observer"
)

func TestMail(t *testing.T) {
	sendEmail := &observer.Email{
		Name: "send_email",
	}

	item := &observer.Item{
		Name: "new_item",
	}
	item.Register(sendEmail)

	// Something happens and we update the item
	item.Update()

	// And we expect the observer to be triggered by that event
	now := time.Now()
	if sendEmail.LastExecution.Format("2006/01/02") != now.Format("2006/01/02") {
		t.Fatal("The observer was not triggered")
	}
}
